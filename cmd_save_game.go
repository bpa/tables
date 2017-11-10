package main

import (
	"encoding/json"
	"errors"

	uuid "github.com/satori/go.uuid"
)

type SaveGameMessage struct {
	Game Game `json:"game"`
}

func SaveGame(c *Client, msg []byte) error {
	var cmd SaveGameMessage
	err := json.Unmarshal(msg, &cmd)
	if err != nil {
		return err
	}

	if cmd.Game.Name == "" {
		return errors.New("game.name is required")
	}

	if cmd.Game.Min == 0 {
		return errors.New("game.min is required")
	}

	if cmd.Game.Max == 0 {
		return errors.New("game.max is required")
	}

	_, ok := Games[cmd.Game.Id]
	if !ok {
		cmd.Game.Id = uuid.NewV4().String()
	}

	Games[cmd.Game.Id] = cmd.Game
	saveState()
	g, _ := json.Marshal(GamesMessage{"games", Games})
	c.hub.broadcast <- g
	return nil
}
