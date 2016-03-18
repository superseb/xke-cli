package xke

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Client struct {
	Token       string
	httpClient  *http.Client
	ListURL     string
	SessionsURL string
}

func NewClient(token string) *Client {
	return &Client{
		Token:       token,
		httpClient:  &http.Client{},
		ListURL:     "https://xke-nxt.appspot.com/api/xke/?format=json&ordering=ddate",
		SessionsURL: "https://xke-nxt.appspot.com/api/session/?format=json&xke=2016-03-18",
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
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error reading [%s] got response [%v]", url, resp.StatusCode)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
