package winestro

import (
	"encoding/xml"
	"fmt"
	"regexp"
	"strings"
)

type Customer struct {
	ID               int32   `xml:"adr_id"`
	Number           string  `xml:"adr_nr"`
	FirstName        string  `xml:"adr_vorname"`
	LastName         string  `xml:"adr_nachname"`
	Company          string  `xml:"adr_firma"`
	ZIP              string  `xml:"adr_plz"`
	City             string  `xml:"adr_ort"`
	WWW              string  `xml:"adr_www"`
	Email            string  `xml:"adr_email"`
	Street           string  `xml:"adr_str"`
	HouseNum         string  `xml:"adr_str_nr"`
	Country          string  `xml:"adr_land"`
	Landline         string  `xml:"adr_festnetz"`
	Mobile           string  `xml:"adr_mobil"`
	Fax              string  `xml:"adr_fax"`
	Note1            string  `xml:"adr_note1"`
	Note2            string  `xml:"adr_note2"`
	Note3            string  `xml:"adr_note3"`
	Note4            string  `xml:"adr_note4"`
	Discount         float64 `xml:"adr_rabatt"`
	PriceCategory    int     `xml:"adr_id_preiskategorie"`
	NewsletterActive bool    `xml:"adr_newsletter_aktiv"`
	TaxType          int     `xml:"adr_kunden_mwst"`
	Salutation       string  `xml:"adr_anrede"`
	SalutationType   int     `xml:"adr_anredenart"`
	PaymentType      int     `xml:"adr_id_zahlungsart"`
}

func (c Customer) FullName() string {
	if c.FirstName == "" && c.LastName == "" {
		return c.Company
	}
	return fmt.Sprintf("%s %s", c.FirstName, c.LastName)
}

func (c Customer) FullSalutation() string {
	re := regexp.MustCompile(`,\s*$`)
	trimSaluation := strings.TrimSpace(re.ReplaceAllString(c.Salutation, ""))
	switch c.SalutationType {
	case 0:
		return fmt.Sprintf("%s %s", trimSaluation, strings.TrimSpace(c.FirstName))
	case 1:
		return fmt.Sprintf("%s %s", trimSaluation, strings.TrimSpace(c.LastName))
	default:
		return trimSaluation
	}
}

type CustomersResp struct {
	XMLName   xml.Name   `xml:"items"`
	Customers []Customer `xml:"item"`
}

func (c *Client) FetchCustomersForGroup(groupID int) ([]Customer, error) {
	params := map[string]string{
		"id_grp": fmt.Sprintf("%d", groupID),
	}
	var customers CustomersResp
	err := c.do("getKundenGruppe", params, &customers)
	return customers.Customers, err
}
