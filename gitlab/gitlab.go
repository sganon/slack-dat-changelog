package gitlab

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
)

type Client struct {
	accessToken string
	client      *http.Client
}

func New(accessToken string) *Client {
	return &Client{
		accessToken: accessToken,
		client:      http.DefaultClient,
	}
}

func (c Client) GetRawFile(repoURL, filename string) (content []byte, err error) {
	var body io.Reader
	req, err := http.NewRequest("GET", repoURL+"/raw/master/"+filename, body)
	if err != nil {
		return content, fmt.Errorf("error in GetRawFile: %v", err)
	}
	req.Header.Set("PRIVATE-TOKEN", c.accessToken)
	res, err := c.client.Do(req)
	if err != nil {
		return content, fmt.Errorf("error in GetRawFile: %v", err)
	}
	if res.StatusCode != http.StatusOK {
		return content, fmt.Errorf("gitlab api returned status %d instead of 200", res.StatusCode)
	}
	content, err = ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	return content, err
}
