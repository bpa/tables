package main

import (
	"fmt"
	"time"
)

type Notifier interface {
	notify(p *Player, message string)
}

func NotifyNewTable(table *Table, player *Player) {
	msg := fmt.Sprintf("%s created a table for %s at %s (%s)",
		player.FullName, table.Game.Name, table.Start.Local().Format(time.Kitchen), SiteUrl)
	for name := range Notifications {
		notifications := Notifications[name]
		notifier := Notifiers[name]
		if notifier != nil {
			for p := range notifications {
				notifier.notify(&notifications[p], msg)
			}
		}
	}
}
