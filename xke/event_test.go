package xke_test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/rchrd/xke-cli/xke"
)

var allEvents = []xke.Event{
	{time.Now().AddDate(0, 0, -14).Format("2006-01-02"), "XKE"},
	{time.Now().Format("2006-01-02"), "Innovation Day"},
	{time.Now().AddDate(0, 0, 14).Format("2006-01-02"), "XKE"},
}

func TestAllEvents(t *testing.T) {
	token := "MySecretToken"

	eventsIn, _ := json.Marshal(allEvents)

	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
		fmt.Fprintln(w, string(eventsIn))
	}
	server := httptest.NewServer(http.HandlerFunc(f))
	defer server.Close()

	c := xke.NewClient(token)
	c.ListURL = server.URL
	eventsOut, _ := c.AllEvents()

	if len(eventsOut) != 3 {
		t.Errorf("Incorrect number of events: expected %v got %v", 3, len(eventsOut))
	}
}
