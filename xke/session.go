package xke

import (
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"

	"github.com/fatih/color"
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
	bold := color.New(color.Bold).SprintFunc()
	t := s.Title
	if len(t) > 65 {
		t = s.Title[:62] + "..."
	}
	p := s.Presenter
	if len(p) > 30 {
		p = s.Presenter[:27] + "..."
	}
	l := s.Location.Name
	if len(l) > 25 {
		l = s.Location.Name[:22] + "..."
	}
	return fmt.Sprintf("%4d  %s  %-75s  %-30s  %-25s", s.ID, s.StartTime[:5], bold(t), p, l)
}

func (s Session) PrintDetails() string {
	label := color.New(color.Bold).SprintFunc()
	details := fmt.Sprintf("%-20s  %s\n", label("Title"), s.Title)
	details += fmt.Sprintf("%-20s  %s\n", label("Presenter"), s.Presenter)
	details += fmt.Sprintf("\n")
	details += fmt.Sprintf("%-20s  %s - %s\n", label("Timeslot"), s.StartTime[:5], s.EndTime[:5])
	details += fmt.Sprintf("%-20s  %s\n", label("Location"), s.Location.Name)
	details += fmt.Sprintf("%-20s  %s\n", label("Type"), s.SessionType)
	details += fmt.Sprintf("%-20s  %s\n", label("Unit"), s.Unit)
	details += fmt.Sprintf("\n")
	details += fmt.Sprintf("%s\n", label("Summary"))
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

	b := make(chan int)
	for i := range sessions {
		go c.addLocation(&sessions[i], b)
	}
	for range sessions {
		<-b // wait for each task to complete
	}
	// all done
	return sessions, nil
}

func (c *Client) Session(id int) (Session, error) {
	u := c.SessionURL
	u.Path = u.Path + strconv.Itoa(id) + "/"
	content, _ := c.getContent(u)
	s, _ := unmarshalSession(content)
	b := make(chan int)
	go c.addLocation(&s, b)
	<-b
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

func (c *Client) addLocation(s *Session, b chan int) {
	u, _ := url.Parse(s.LocationURL)
	s.Location, _ = c.LocationByURL(u)
	b <- 1 // signal that this call is complete
}
