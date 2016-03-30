package xke

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Client struct {
	Token       string
	httpClient  *http.Client
	ListURL     *url.URL
	SessionURL  *url.URL
	LocationURL *url.URL
}

func NewClient(token string) (*Client, error) {
	if token == "" {
		return nil, errors.New("Authentication token not found. Please set it via flag or environment variable.")
	}
	c := &Client{
		Token:      token,
		httpClient: &http.Client{},
	}
	c.ListURL, _ = url.Parse("https://xke-nxt.appspot.com/api/xke/")
	c.SessionURL, _ = url.Parse("https://xke-nxt.appspot.com/api/session/")
	c.LocationURL, _ = url.Parse("https://xke-nxt.appspot.com/api/location/")
	return c, nil
}

func (c *Client) getContent(u *url.URL) ([]byte, error) {
	q := u.Query()
	q.Set("format", "json")
	u.RawQuery = q.Encode()
	req, err := http.NewRequest("GET", u.String(), nil)
	req.Header.Add("Authorization", "Token "+c.Token)

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Error reading [%s] got response [%v]", u.String(), resp.StatusCode)
	}

	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	return content, nil
}
