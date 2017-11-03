package tables

import "encoding/json"

type Table struct {
	Name string
	//players []string
	//start   time.Time
	//end     time.Time
}

type TableStruct struct {
	Tables []Table
}

func DecodeTable(d *json.Decoder) (interface{}, error) {
	var t Table
	err := d.Decode(&t)
	return t, err
}

var Tables []Table = make([]Table, 0, 3)

func GetTables() interface{} {
	return TableStruct{Tables: Tables}
}

func AddTable(t interface{}) interface{} {
	table := t.(Table)
	Tables = append(Tables, table)
	return TableStruct{Tables: Tables}
}
