package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"time"
)

type CreateTableMessage struct {
	Game     string    `json:"game"`
	Location string    `json:"location"`
	Start    time.Time `json:"start"`
	Player   Player    `json:"player"`
}

func CreateTable(c *Client, msg []byte) error {
	var cmd CreateTableMessage
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

	if cmd.Player.FullName == "" {
		return errors.New("Player is required")
	}

	game := Games[cmd.Game]
	if game.Name == "" {
		return errors.New(fmt.Sprintf("No game %s available", cmd.Game))
	}

	players := make([]Player, 0, game.Max)
	players = append(players, cmd.Player)

	table := AddNewTable(game, cmd.Location, cmd.Start, players)

	g, _ := json.Marshal(GetTables())
	c.hub.broadcast <- g

	go NotifyNewTable(&table, &cmd.Player)
	return nil
}
