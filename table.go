package main

import (
	"encoding/json"
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
