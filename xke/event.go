package xke

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Date    string `json:"ddate"`
	Comment string `json:"comment"`
	// StartTime   string `json:"start_time"`
	// EndTime     string `json:"end_time"`
	// SessionList string `json:"session_list"`
}

func (e Event) String() string {
	return fmt.Sprintf("%s - %s", e.Date, e.Comment)
}

// AllEvents returns a slice of future and past events
func (c *Client) AllEvents() ([]Event, error) {
	u := c.ListURL
	q := u.Query()
	q.Set("ordering", "ddate")
	u.RawQuery = q.Encode()
	content, _ := c.getContent(u)
	return unmarshalEvents(content)
}

// FutureEvents returns a slice of future events
func (c *Client) FutureEvents() ([]Event, error) {
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
	return unmarshalEvents(content)
}

// NextEvent returns the next event based on date
func (c *Client) NextEvent() (Event, error) {
	events, err := c.FutureEvents()
	return events[0], err
}

func unmarshalEvents(content []byte) ([]Event, error) {
	var events []Event
	err := json.Unmarshal(content, &events)
	return events, err
}
