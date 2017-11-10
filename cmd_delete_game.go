package main

import (
	"encoding/json"
	"errors"
)

type DeleteGameMessage struct {
	Game string `json:"game"`
}

func DeleteGame(c *Client, msg []byte) error {
	var cmd DeleteGameMessage
	err := json.Unmarshal(msg, &cmd)
	if err != nil {
		return err
	}

	if cmd.Game == "" {
		return errors.New("Game is required")
	}

	_, ok := Games[cmd.Game]
	if ok {
		delete(Games, cmd.Game)
		saveState()
		g, _ := json.Marshal(GamesMessage{"games", Games})
		c.hub.broadcast <- g
		return nil
	}
	return errors.New("Game does not exist")
}
