package server

import (
	"errors"

	"github.com/bpa/tables/data"
)

type locationMessage struct {
	Location string `json:"location"`
}

func init() {
	commands["create_location"] = createLocation
}

func createLocation(c data.Client, msg obj) error {
	location, _ := msg["location"].(string)
	if location == "" {
		return errors.New("Location is required")
	}

	if err := data.CreateLocation(string(location)); err != nil {
		return err
	}
	c.Broadcast(locationsMessage{"locations", data.GetLocations()})
	return nil
}
