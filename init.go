package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
)

var secret string
var SiteUrl string
var Games map[string]Game
var Locations []string
var Tables []Table
var Notifiers = make(map[string]Notifier)
var Notifications map[string][]Player

type config struct {
	Secret         string
	SiteUrl        string
	Authentication map[string]string
	Notifications  map[string]map[string]string
}

type memory struct {
	Games         map[string]Game
	Locations     []string
	Tables        []Table
	Notifications map[string][]Player
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

	if conf.Authentication["type"] == "LDAP" {
		commands["login"] = LDAPLogin
		LDAPInit(conf.Authentication)
	}
	for k := range conf.Notifications {
		notifier := conf.Notifications[k]
		if notifier["type"] == "http" {
			note, err := NewNotifyHttp(notifier)
			if err != nil {
				log.Fatal(fmt.Sprintf("Error in config.json/notifications/%s: %s", k, err.Error()))
			}
			Notifiers[k] = note
		}
	}
	secret = conf.Secret
	SiteUrl = conf.SiteUrl
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
			Notifications = mem.Notifications
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
	if Notifications == nil {
		Notifications = make(map[string][]Player)
	}
}

func saveState() {
	f, err := os.OpenFile("state.json", os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	if err == nil {
		defer f.Close()

		enc := json.NewEncoder(f)
		err = enc.Encode(memory{
			Games:         Games,
			Locations:     Locations,
			Tables:        Tables,
			Notifications: Notifications,
		})
		if err != nil {
			fmt.Printf("Error writing state.json: %s\n", err)
		}
	}
}
