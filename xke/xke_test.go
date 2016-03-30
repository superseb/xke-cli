package xke_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
	"time"

	"github.com/rchrd/xke-cli/xke"
)

var (
	allEvents = []xke.XKE{
		{time.Now().AddDate(0, 0, -14).Format("2006-01-02"), "XKE"},
		{time.Now().Format("2006-01-02"), "Innovation Day"},
		{time.Now().AddDate(0, 0, 14).Format("2006-01-02"), "XKE"},
	}
	futureEvents = []xke.XKE{
		{time.Now().Format("2006-01-02"), "Innovation Day"},
		{time.Now().AddDate(0, 0, 14).Format("2006-01-02"), "XKE"},
	}
)

func TestAllEvents(t *testing.T) {
	eventsIn, _ := json.Marshal(allEvents)

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(eventsIn))
	}
	server := httptest.NewServer(http.HandlerFunc(f))
	defer server.Close()

	c, _ := xke.NewClient("token")
	c.ListURL, _ = url.Parse(server.URL)
	eventsOut, _ := c.AllXKEs()

	if len(eventsOut) != 3 {
		t.Errorf("Incorrect number of events: expected %v got %v", 3, len(eventsOut))
	}
}

func TestFutureEvents(t *testing.T) {
	eventsIn, _ := json.Marshal(futureEvents)

	f := func(w http.ResponseWriter, r *http.Request) {
		if !strings.Contains(r.RequestURI, time.Now().Format("2006-01-02")) {
			t.Error("Query string does not contain current date")
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(eventsIn))
	}
	server := httptest.NewServer(http.HandlerFunc(f))
	defer server.Close()

	c, _ := xke.NewClient("token")
	c.ListURL, _ = url.Parse(server.URL)
	eventsOut, _ := c.FutureXKEs()

	if len(eventsOut) != 2 {
		t.Errorf("Incorrect number of events: expected %v got %v", 2, len(eventsOut))
	}
}
