package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{}

type Client interface {
	setPlayer(*Player)
	send(interface{}) error
	host() string
}

type WSClient struct {
	conn       *websocket.Conn
	player     Player
	output     chan []byte
	remoteHost string
	hub        *WSHub
}

func (c *WSClient) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		handleMessage(c, message)
	}
}

func (c WSClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.output:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteMessage(websocket.TextMessage, message)
			if err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c WSClient) send(v interface{}) error {
	g, err := json.Marshal(v)
	if err != nil {
		return err
	}
	c.output <- g
	return nil
}

func (c WSClient) host() string {
	return c.remoteHost
}

func (c WSClient) setPlayer(p *Player) {
	c.player = *p
}

func serveWs(hub *WSHub, w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	client := &WSClient{
		conn:       conn,
		output:     make(chan []byte, 256),
		remoteHost: r.Header.Get("X-Forwarded-For"),
		hub:        hub,
	}
	if client.remoteHost == "" {
		addr := conn.RemoteAddr().String()
		i := strings.LastIndexByte(addr, ':')
		if i != -1 {
			addr = addr[:i]
		}
		client.remoteHost = addr
	}
	client.hub.register <- client

	go client.writePump()
	go client.readPump()

	client.send(GetTables())
}
