package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type CreateGameMessage struct {
	Game     string    `json:"game"`
	Location string    `json:"location"`
	Start    time.Time `json:"start"`
}

func CreateGame(c *Client, msg []byte) error {
	var cmd CreateGameMessage
	err := json.Unmarshal(msg, &cmd)
	if err != nil {
		return err
	}

	if cmd.Game == "" {
		return errors.New("Game is required")
	}

	if cmd.Location == "" {
		return errors.New("Location is required")
	}

	if cmd.Start.IsZero() {
		return errors.New("Start is required")
	}

	game := Games[cmd.Game]
	if game.Name == "" {
		return errors.New(fmt.Sprintf("No game %s available", cmd.Game))
	}

	Tables = append(Tables, Table{
		Game:     game,
		Location: cmd.Location,
		Start:    cmd.Start,
	})

	g, _ := json.Marshal(GetTables())
	c.hub.broadcast <- g
	return nil
}
