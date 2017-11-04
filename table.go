package tables

import (
	"encoding/json"
	"time"
)

type Table struct {
	Name     string    `json:"name"`
	Players  []Player  `json:"players"`
	Min      int       `json:"min"`
	Max      int       `json:"max"`
	Location string    `json:"location"`
	Start    time.Time `json:"start"`
}

type TableStruct struct {
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
	return TableStruct{Cmd: "tables", Tables: Tables}
}

func AddTable(t interface{}) interface{} {
	table := t.(Table)
	Tables = append(Tables, table)
	return TableStruct{Tables: Tables}
}
