package server

import (
	"errors"
	"time"

	"github.com/bpa/tables/data"
	"github.com/bpa/tables/notify"
	"github.com/mitchellh/mapstructure"
)

func init() {
	commands["create_table"] = authenticated(createTable)
}

type createTableMessage struct {
	Game     string
	Location string
	Start    string
}

func createTable(c data.Client, msg obj) error {
	var cmd createTableMessage
	if err := mapstructure.Decode(msg, &cmd); err != nil {
		return errors.New("Bad message")
	}

	if cmd.Game == "" {
		return errors.New("Game is required")
	}

	if cmd.Location == "" {
		return errors.New("Location is required")
	}

	if cmd.Start == "" {
		return errors.New("Start is required")
	}

	start, err := time.Parse(time.RFC3339, cmd.Start)
	if err != nil {
		start, err = time.Parse("2006-01-02T15:04:05-0700", cmd.Start)
		if err != nil {
			return err
		}
	}
	table, err := data.CreateTable(cmd.Game, cmd.Location, start, c.GetPlayer())
	if err != nil {
		return err
	}

	go notify.NotifyNewTable(table, c.GetPlayer())

	return hub.Broadcast(tablesMessage{"tables", data.GetTables()})
}
