package main

import (
	"encoding/json"
	"errors"
)

type CreateLocationMessage struct {
	Location string `json:"location"`
}

func CreateLocation(c Client, msg []byte) error {
	var cmd CreateLocationMessage
	err := json.Unmarshal(msg, &cmd)
	if err != nil {
		return err
	}

	if cmd.Location == "" {
		return errors.New("Location is required")
	}

	for i := range Locations {
		if cmd.Location == Locations[i] {
			return errors.New("Location already exists")
		}
	}

	Locations = append(Locations, cmd.Location)

	saveState()

	return hub.Broadcast(LocationsMessage{"locations", Locations})
}
