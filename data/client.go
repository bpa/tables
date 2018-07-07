package data

type Client interface {
	Broadcast(interface{})
	Error(string, ...interface{})
	GetPlayer() *Player
	Host() string
	Send(interface{})
	SetPlayer(*Player)
}
