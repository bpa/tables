package main

import (
	"encoding/json"
	"log"
	"os"
)

var secret string
var Games []Game
var Locations []string
var Tables []Table

type config struct {
	Secret string
}

type memory struct {
	Games     []Game
	Locations []string
	Tables    []Table
}

func LoadStartupFiles() {
	readConfig()
	readState()
}

func readConfig() {
	f, err := os.Open("config.json")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	var conf config
	dec := json.NewDecoder(f)
	err = dec.Decode(&conf)
	if err != nil {
		log.Fatalf("Can't read config.json: %s", err)
	}

	secret = conf.Secret
}

func readState() {
	f, err := os.Open("state.json")
	if err == nil {
		defer f.Close()

		var mem memory
		dec := json.NewDecoder(f)
		err = dec.Decode(&mem)
		if err == nil {
			Tables = mem.Tables
			Games = mem.Games
			Locations = mem.Locations
		}
	}

	if Tables == nil {
		Tables = make([]Table, 0, 3)
	}
	if Games == nil {
		Games = make([]Game, 0, 3)
	}
	if Locations == nil {
		Locations = make([]string, 0, 3)
	}
}
