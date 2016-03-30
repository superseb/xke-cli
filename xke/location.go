package xke

import (
	"encoding/json"
	"fmt"
	"net/url"
)

type Location struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	MaxCapacity int    `json:"max_capacity"`
	Address     string `json:"address"`
}

func (l Location) String() string {
	return fmt.Sprintf("%2v - %s - %v - %s", l.ID, l.Name, l.MaxCapacity, l.Address)
}

func (c *Client) Locations() ([]Location, error) {
	content, _ := c.getContent(c.LocationURL)
	return unmarshalLocations(content)
}

func (c *Client) LocationByID(id int) (Location, error) {
	u := c.LocationURL
	u.Path += string(id)
	return c.LocationByURL(u)
}

func (c *Client) LocationByURL(u *url.URL) (Location, error) {
	content, _ := c.getContent(u)
	locations, err := unmarshalLocation(content)
	return locations, err
}

func unmarshalLocations(content []byte) ([]Location, error) {
	var locations []Location
	err := json.Unmarshal(content, &locations)
	return locations, err
}

func unmarshalLocation(content []byte) (Location, error) {
	var location Location
	err := json.Unmarshal(content, &location)
	return location, err
}
