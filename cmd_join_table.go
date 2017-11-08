package main

import (
	"encoding/json"
)

type JoinTableMessage struct {
	Id     string `json:"id"`
	Player Player `json:"player"`
}

func JoinTable(c *Client, msg []byte) error {
	var cmd JoinTableMessage
	err := json.Unmarshal(msg, &cmd)
	if err != nil {
		return err
	}

	err = AddPlayerToTable(&cmd.Player, cmd.Id)
	if err != nil {
		return err
	}

	g, _ := json.Marshal(GetTables())
	c.hub.broadcast <- g
	return nil
}
