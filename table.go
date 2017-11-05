package tables

import (
	"encoding/json"
	"time"
)

type Table struct {
	Game     Game      `json:"game"`
	Players  []Player  `json:"players"`
	Location string    `json:"location"`
	Start    time.Time `json:"start"`
}

type TableMsg struct {
	Cmd    string  `json:"cmd"`
	Tables []Table `json:"tables"`
}

func DecodeTable(d *json.Decoder) (interface{}, error) {
	var t Table
	err := d.Decode(&t)
	return t, err
}

var Tables []Table = make([]Table, 0, 3)

func GetTables() interface{} {
	return TableMsg{Cmd: "tables", Tables: Tables}
}

func AddTable(t interface{}) interface{} {
	table := t.(Table)
	Tables = append(Tables, table)
	return TableMsg{Tables: Tables}
}
