package data

type Notifier interface {
	NotifyNewTable(table *Table, creator *Player, notify *[]Player)
}
