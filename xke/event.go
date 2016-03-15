package xke

import (
	"encoding/json"
	"fmt"
	"time"
)

type Event struct {
	Date    string `json:"ddate"`
	Comment string
	// StartTime   string `json:"start_time"`
	// EndTime     string `json:"end_time"`
	// SessionList string `json:"session_list"`
}

func (e Event) String() string {
	return fmt.Sprintf("%s - %s", e.Date, e.Comment)
}

func (c *Client) AllEvents() ([]Event, error) {
	content, _ := c.getContent(ListURL)
	var events []Event
	err := json.Unmarshal(content, &events)
	return events, err
}

func (c *Client) FutureEvents() ([]Event, error) {
	content, _ := c.getContent(ListURL + "&ddate__gte=" + time.Now().Format("2006-01-02"))
	var events []Event
	err := json.Unmarshal(content, &events)
	return events, err
}
