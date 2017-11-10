package main

import (
	"encoding/json"
	"errors"
)

func DeleteLocation(c *Client, msg []byte) error {
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
			Locations = append(Locations[:i], Locations[i+1:]...)
			saveState()
			g, _ := json.Marshal(LocationsMessage{"locations", Locations})
			c.hub.broadcast <- g
			return nil
		}
	}
	return errors.New("Location does not exist")
}
