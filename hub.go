package main

import "encoding/json"

type Hub interface {
	Broadcast(interface{}) error
}

//This class taken almost directly from https://github.com/gorilla/websocket/blob/master/examples/chat/hub.go
type WSHub struct {
	clients    map[*WSClient]bool
	broadcast  chan []byte
	register   chan *WSClient
	unregister chan *WSClient
}

func newWSHub() *WSHub {
	return &WSHub{
		broadcast:  make(chan []byte),
		register:   make(chan *WSClient),
		unregister: make(chan *WSClient),
		clients:    make(map[*WSClient]bool),
	}
}

func (h *WSHub) Broadcast(v interface{}) error {
	g, err := json.Marshal(v)
	if err != nil {
		return err
	}
	h.broadcast <- g
	return nil
}

func (h *WSHub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				close(client.output)
				delete(h.clients, client)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.output <- message:
				default:
					close(client.output)
					delete(h.clients, client)
				}
			}
		}
	}
}
