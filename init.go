package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var secret string
var Games map[string]Game
var Locations []string
var Tables []Table

type config struct {
	Secret string
}

type memory struct {
	Games     map[string]Game
	Locations []string
	Tables    []Table
}

func loadStartupFiles() {
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
		} else {
			fmt.Printf("Error reading state.json: %s\n", err)
		}
	}

	if Tables == nil {
		Tables = make([]Table, 0, 3)
	}
	if Games == nil {
		Games = make(map[string]Game)
	}
	if Locations == nil {
		Locations = make([]string, 0, 3)
	}
}

func saveState() {
	f, err := os.OpenFile("state.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err == nil {
		defer f.Close()

		enc := json.NewEncoder(f)
		err = enc.Encode(memory{
			Games:     Games,
			Locations: Locations,
			Tables:    Tables,
		})
		if err != nil {
			fmt.Printf("Error writing state.json: %s\n", err)
		}
	}
}
