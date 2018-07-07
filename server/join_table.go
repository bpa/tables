package server

import (
	"errors"

	"github.com/bpa/tables/data"
)

func init() {
	commands["join_table"] = authenticated(joinTable)
}

func joinTable(c data.Client, msg obj) error {
	id, _ := msg["id"].(string)
	if id == "" {
		return errors.New("Missing 'id'")
	}

	player := c.GetPlayer()
	if player == nil {
		return errors.New("Not logged in")
	}

	if err := data.JoinTable(player, id); err != nil {
		return err
	}

	return hub.Broadcast(tablesMessage{"tables", data.GetTables()})
}
