package xke

import (
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

var token = "MyToken"

func TestGetContentAuthorization(t *testing.T) {
	f := func(w http.ResponseWriter, r *http.Request) {
		if strings.Join(r.Header["Authorization"], "") != "Token "+token {
			t.Error("Authorization token incorrect")
		}
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "application/json")
	}
	server := httptest.NewServer(http.HandlerFunc(f))
	defer server.Close()

	client := NewClient(token)
	u, _ := url.Parse(server.URL)
	client.getContent(u)
}

func TestGetContentNotStatusOK(t *testing.T) {
	f := func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(403)
		w.Header().Set("Content-Type", "application/json")
	}
	server := httptest.NewServer(http.HandlerFunc(f))
	defer server.Close()

	client := NewClient(token)
	u, _ := url.Parse(server.URL)
	_, err := client.getContent(u)
	if err == nil {
		t.Error("HTTP response not equal to 200 should raise an error")
	}
}
