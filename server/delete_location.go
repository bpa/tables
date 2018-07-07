package server

import (
	"errors"

	"github.com/bpa/tables/data"
)

func init() {
	commands["delete_location"] = deleteLocation
}

func deleteLocation(c data.Client, msg obj) error {
	location, _ := msg["location"].(string)
	if location == "" {
		return errors.New("Location is required")
	}
	if err := data.DeleteLocation(location); err != nil {
		return err
	}
	return hub.Broadcast(locationsMessage{"locations", data.GetLocations()})
}
