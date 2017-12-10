package main

import (
	"encoding/json"
)

type LeaveTableMessage struct {
	Id     string `json:"id"`
	Player Player `json:"player"`
}

func LeaveTable(c Client, msg []byte) error {
	var cmd LeaveTableMessage
	err := json.Unmarshal(msg, &cmd)
	if err != nil {
		return err
	}

	err = DeletePlayerFromTable(cmd.Player.Id, cmd.Id)
	if err != nil {
		return err
	}

	return hub.Broadcast(GetTables())
}
