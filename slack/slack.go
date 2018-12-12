package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Payload struct {
	Attachments []Attachment `json:"attachments"`
}

type Attachment struct {
	Fallback string  `json:"fallback"`
	Pretext  string  `json:"pretext"`
	Color    string  `json:"color"`
	Fields   []Field `json:"fields"`
}

type Field struct {
	Title string `json:"title"`
	Value string `json:"value"`
	Short bool   `json:"bool"`
}

type Client struct {
	webHookURL string
	channel    string
	client     *http.Client
}

func New(webHookURL, channel string) *Client {
	return &Client{
		webHookURL: webHookURL,
		channel:    channel,
		client:     http.DefaultClient,
	}
}

func (c Client) SendMessage(payload Payload) (err error) {
	b, err := json.Marshal(payload)
	if err != nil {
		return fmt.Errorf("error in send messag: %v", err)
	}
	req, err := http.NewRequest("POST", c.webHookURL, bytes.NewBuffer(b))
	if err != nil {
		return fmt.Errorf("error in send messag: %v", err)
	}
	res, err := c.client.Do(req)
	if err != nil {
		return fmt.Errorf("error in send messag: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("slack webhook returned %d instead of 200", res.StatusCode)
	}
	return err
}
