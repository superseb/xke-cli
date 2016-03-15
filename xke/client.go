package xke

import (
	"io/ioutil"
	"net/http"
)

// Endpoints of the XKE API
const (
	ListURL     = "https://xke-nxt.appspot.com/api/xke/?format=json"
	SessionsURL = "https://xke-nxt.appspot.com/api/session/?format=json&xke=2016-03-18"
)

type Client struct {
	Token      string
	httpClient *http.Client
}

func NewClient(token string) *Client {
	return &Client{
		Token:      token,
		httpClient: &http.Client{},
	}
}

func (c *Client) getContent(url string) ([]byte, error) {
	req, err := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Token "+c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
