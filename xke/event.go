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
	content, _ := c.getContent(c.ListURL)
	return unmarshalEvents(content)
}

// FutureEvents returns a slice of future events
func (c *Client) FutureEvents() ([]Event, error) {
	content, _ := c.getContent(c.ListURL + "&ddate__gte=" + time.Now().Format("2006-01-02"))
	return unmarshalEvents(content)
}

func unmarshalEvents(content []byte) ([]Event, error) {
	var events []Event
	err := json.Unmarshal(content, &events)
	return events, err
}
