package xke

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
)

type Session struct {
	ID          int    `json:"id"`
	SessionType string `json:"session_type"`
	Title       string `json:"title"`
	Presenter   string `json:"presenter"`
	XKE         string `json:"xke"`
	StartTime   string `json:"start_time"`
	EndTime     string `json:"end_time"`
	Unit        string `json:"unit"`
	LocationURL string `json:"location"`
	Goal        string `json:"goal"`
	Summary     string `json:"summary"`
	Preparation string `json:"preparation"`
	Location
}

func (s Session) PrintSummary() string {
	return fmt.Sprintf("%4d - %s - %s [%s] - %s", s.ID, s.StartTime, s.Title, s.Presenter, s.Location.Name)
}

func (s Session) PrintDetails() string {
	details := fmt.Sprintf("%s - %s - %s - %s - %s\n\n", s.StartTime, s.EndTime, s.SessionType, s.Unit, s.Location.Name)
	details += fmt.Sprintf("%s [%s]\n\n", s.Title, s.Presenter)
	details += fmt.Sprintf("%s", s.Summary)
	return details
}

func (c *Client) Sessions(date string) ([]Session, error) {
	u := c.SessionURL
	q := u.Query()
	q.Set("ordering", "start_time")
	q.Set("xke", date)
	u.RawQuery = q.Encode()
	content, _ := c.getContent(u)
	sessions, _ := unmarshalSessions(content)
	for i := range sessions {
		c.addLocation(&sessions[i])
	}
	return sessions, nil
}

func (c *Client) Session(id int) (Session, error) {
	u := c.SessionURL
	u.Path = u.Path + strconv.Itoa(id) + "/"
	content, _ := c.getContent(u)
	s, _ := unmarshalSession(content)
	c.addLocation(&s)
	return s, nil
}

func unmarshalSessions(content []byte) ([]Session, error) {
	var sessions []Session
	err := json.Unmarshal(content, &sessions)
	return sessions, err
}

func unmarshalSession(content []byte) (Session, error) {
	var session Session
	err := json.Unmarshal(content, &session)
	return session, err
}

func (c *Client) addLocation(s *Session) {
	u, _ := url.Parse(s.LocationURL)
	s.Location, _ = c.LocationByURL(u)
}
