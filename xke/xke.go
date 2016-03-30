package xke

import (
	"encoding/json"
	"fmt"
	"time"
)

type XKE struct {
	Date    string `json:"ddate"`
	Comment string `json:"comment"`
	// StartTime   string `json:"start_time"`
	// EndTime     string `json:"end_time"`
	// SessionList string `json:"session_list"`
}

func (e XKE) String() string {
	return fmt.Sprintf("%s - %s", e.Date, e.Comment)
}

// AllXKEs returns a slice of future and past xkes
func (c *Client) AllXKEs() ([]XKE, error) {
	u := c.ListURL
	q := u.Query()
	q.Set("ordering", "ddate")
	u.RawQuery = q.Encode()
	content, _ := c.getContent(u)
	return unmarshalXKEs(content)
}

// FutureXKEs returns a slice of future xkes
func (c *Client) FutureXKEs() ([]XKE, error) {
	u := c.ListURL
	q := u.Query()
	q.Set("ordering", "ddate")
	q.Set("ddate__gte", time.Now().Format("2006-01-02"))
	u.RawQuery = q.Encode()
	content, err := c.getContent(u)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return unmarshalXKEs(content)
}

// NextXKE returns the next xke based on date
func (c *Client) NextXKE() (XKE, error) {
	xkes, err := c.FutureXKEs()
	return xkes[0], err
}

func unmarshalXKEs(content []byte) ([]XKE, error) {
	var xkes []XKE
	err := json.Unmarshal(content, &xkes)
	return xkes, err
}
