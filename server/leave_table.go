package server

import (
	"errors"

	"github.com/bpa/tables/data"
)

func init() {
	commands["leave_table"] = authenticated(leaveTable)
}

func leaveTable(c data.Client, msg obj) error {
	id, _ := msg["id"].(string)
	if id == "" {
		return errors.New("Missing 'id'")
	}
	if err := data.LeaveTable(c.GetPlayer(), id); err != nil {
		return err
	}
	return hub.Broadcast(tablesMessage{"tables", data.GetTables()})
}
