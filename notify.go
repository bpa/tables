package main

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
