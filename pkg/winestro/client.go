package winestro

import (
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"time"
)

const host = "https://weinstore.net/xml/v21.0/wbo-API.php"

type Client struct {
	httpClient *http.Client
	uid        int
	user       string
	code       string
	shopID     int
}

func NewClient(uid int, user string, code string, shopID int) *Client {
	return &Client{
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
		uid:    uid,
		user:   user,
		code:   code,
		shopID: shopID,
	}
}

func (c *Client) do(action string, params map[string]string, v interface{}) error {
	req, err := http.NewRequest("POST", host, nil)
	if err != nil {
		return err
	}
	q := req.URL.Query()
	q.Add("UID", fmt.Sprint(c.uid))
	q.Add("apiUSER", c.user)
	q.Add("apiCODE", c.code)
	q.Add("apiShopID", fmt.Sprint(c.shopID))
	q.Add("apiACTION", action)
	q.Add("output", "xml")
	for key, value := range params {
		q.Add(key, value)
	}
	req.URL.RawQuery = q.Encode()
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	return decode(resp.Body, v)
}

func decode(r io.Reader, v interface{}) error {
	return xml.NewDecoder(r).Decode(v)
}
