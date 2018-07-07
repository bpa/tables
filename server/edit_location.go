package server

import (
	"errors"

	"github.com/bpa/tables/data"
)

func init() {
	commands["edit_location"] = editLocation
}

type editLocationMessage struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func editLocation(c data.Client, msg obj) error {
	from, _ := msg["from"].(string)
	to, _ := msg["to"].(string)

	if from == "" {
		return errors.New("from is required")
	}

	if to == "" {
		return errors.New("to is required")
	}

	if err := data.UpdateLocation(from, to); err != nil {
		return err
	}
	return hub.Broadcast(locationsMessage{"locations", data.GetLocations()})
}
