package sync

import (
	"context"
	"encoding/json"
	"log"
	"regexp"
	"strings"

	"github.com/nico0302/winmonk/pkg/listmonk"
)

type customerAttribs struct {
	Number         string `json:"number"`
	FirstName      string `json:"first_name"`
	LastName       string `json:"last_name"`
	Company        string `json:"company"`
	ZIP            string `json:"zip"`
	City           string `json:"city"`
	Country        string `json:"country"`
	Salutation     string `json:"salutation"`
	SalutationType int    `json:"salutation_type"`
	FullSalutation string `json:"full_salutation"`
}

func (c *Client) ImportContacts(ctx context.Context, winGroupID int, listID int) error {
	customers, err := c.winestro.FetchCustomersForGroup(winGroupID)
	if err != nil {
		return err
	}

	r := regexp.MustCompile(`^[^@]+@[^@]+\.[^@]+$`)

	count := 0

	for _, customer := range customers {
		if customer.Email == "" || !r.MatchString(customer.Email) {
			log.Default().Printf("Invalid Email: %s (%s) - %s", customer.FullName(), customer.Number, customer.Email)
			continue
		}

		attribs := customerAttribs{
			Number:         customer.Number,
			FirstName:      customer.FirstName,
			LastName:       customer.LastName,
			Company:        customer.Company,
			ZIP:            customer.ZIP,
			City:           customer.City,
			Country:        customer.Country,
			Salutation:     customer.Salutation,
			SalutationType: customer.SalutationType,
			FullSalutation: customer.FullSalutation(),
		}
		attribsJSON, err := json.Marshal(attribs)
		if err != nil {
			return err
		}

		_, err = c.listmonk.SetSubscriber(ctx, listmonk.SetSubscriberParams{
			ID:      customer.ID,
			Email:   customer.Email,
			Name:    customer.FullName(),
			Attribs: attribsJSON,
		})
		if err != nil {
			if strings.Contains(err.Error(), "subscribers_email_key") || strings.Contains(err.Error(), "idx_subs_email") {
				log.Default().Printf("Duplicate Email: %s (%s) - %s", customer.FullName(), customer.Number, customer.Email)

				// Check if we need to unsubscribe the original email holder
				if customer.NewsletterActive {
					continue
				}
				sub, err := c.listmonk.FindSubscriberByEmail(ctx, customer.Email)
				if err != nil {
					continue
				}
				_, err = c.listmonk.SubscribeList(ctx, listmonk.SubscribeListParams{
					SubscriberID: sub.ID,
					ListID:       int32(listID),
					Status:       listmonk.SubscriptionStatusUnsubscribed,
				})
				if err != nil {
					continue
				}

				continue
			}
			log.Default().Printf("Error: %s (%s) - %s: %s", customer.FullName(), customer.Number, customer.Email, err)
			return err
		}

		status := listmonk.SubscriptionStatusConfirmed
		if !customer.NewsletterActive {
			status = listmonk.SubscriptionStatusUnsubscribed
		}

		// Subscribe to list
		c.listmonk.SubscribeList(ctx, listmonk.SubscribeListParams{
			SubscriberID: customer.ID,
			ListID:       int32(listID),
			Status:       status,
		})

		count++
	}

	log.Default().Printf("Imported %d contacts", count)

	return nil
}
