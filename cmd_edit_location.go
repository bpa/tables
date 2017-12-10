package main

import (
	"encoding/json"
	"errors"
)

type EditLocationMessage struct {
	From string `json:"from"`
	To   string `json:"to"`
}

func EditLocation(c Client, msg []byte) error {
	var cmd EditLocationMessage
	err := json.Unmarshal(msg, &cmd)
	if err != nil {
		return err
	}

	if cmd.From == "" {
		return errors.New("From is required")
	}

	if cmd.To == "" {
		return errors.New("To is required")
	}

	for i := range Locations {
		if cmd.From == Locations[i] {
			Locations[i] = cmd.To
			saveState()
			return hub.Broadcast(LocationsMessage{"locations", Locations})
		}
	}
	return errors.New("Location does not exist")
}
