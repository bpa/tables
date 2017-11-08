package main

import (
	"encoding/json"
	"errors"
	"time"

	uuid "github.com/satori/go.uuid"
)

type Table struct {
	Game     Game      `json:"game"`
	Players  []Player  `json:"players"`
	Location string    `json:"location"`
	Start    time.Time `json:"start"`
	Id       string    `json:"id"`
}

type TablesMessage struct {
	Cmd    string  `json:"cmd"`
	Tables []Table `json:"tables"`
}

func DecodeTable(d *json.Decoder) (interface{}, error) {
	var t Table
	err := d.Decode(&t)
	return t, err
}

func GetTables() TablesMessage {
	return TablesMessage{Cmd: "tables", Tables: Tables}
}

func AddNewTable(game Game, loc string, start time.Time, players []Player) {
	id := uuid.NewV4().String()
	Tables = append(Tables, Table{
		Game:     game,
		Players:  players,
		Location: loc,
		Start:    start,
		Id:       id,
	})
}

func FindTable(id string) (*Table, error) {
	for i := range Tables {
		if Tables[i].Id == id {
			return &Tables[i], nil
		}
	}

	return nil, errors.New("Unknown table")
}

func AddPlayerToTable(player *Player, tableId string) error {
	table, err := FindTable(tableId)
	if err != nil {
		return err
	}

	for i := range table.Players {
		if table.Players[i].Id == player.Id {
			return errors.New("Player already at table")
		}
	}
	table.Players = append(table.Players, *player)
	return nil
}

func DeletePlayerFromTable(playerId string, tableId string) error {
	table, err := FindTable(tableId)
	if err != nil {
		return err
	}

	for i := range table.Players {
		if table.Players[i].Id == playerId {
			table.Players = append(table.Players[:i], table.Players[i+1:]...)
			return nil
		}
	}
	return errors.New("Player isn't at table")
}
