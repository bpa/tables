package server

import (
	"errors"

	"github.com/bpa/tables/data"
)

func init() {
	commands["delete_table"] = deleteTable
}

func deleteTable(c data.Client, msg obj) error {
	table, _ := msg["id"].(string)
	if table == "" {
		return errors.New("Table is required")
	}
	if err := data.DeleteTable(table); err != nil {
		return err
	}
	return hub.Broadcast(tablesMessage{"tables", data.GetTables()})
}
