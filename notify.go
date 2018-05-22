package main

import (
	"errors"
	"fmt"
	"log"
)

type Notifier interface {
	notifyNewTable(table *Table, creator *Player, notify []Player)
}

func NotifyNewTable(table *Table, player *Player) {
	for name := range Notifications {
		notifications := Notifications[name]
		notifier := Notifiers[name]
		if notifier != nil {
			notifier.notifyNewTable(table, player, notifications)
		}
	}
}

func ConfigureNotifications(notifications map[string]ImplEntry) {
	for k := range notifications {
		var err error
		var note Notifier
		notifier := notifications[k]
		switch notifier.Type {
		case "http":
			note, err = NewNotifyHttp(notifier.Config)
		case "smtp":
			note, err = NewNotifySmtp(notifier.Config)
		default:
			err = errors.New(fmt.Sprintf("Unknown type '%s'", notifier.Type))
		}
		if err != nil {
			log.Fatal(fmt.Sprintf("Error in config.json/notifications/%s: %s", k, err.Error()))
		}
		Notifiers[k] = note
	}
}
