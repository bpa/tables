package server

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bpa/tables/data"
	"github.com/gorilla/websocket"
)

const (
	writeWait      = 10 * time.Second
	pongWait       = 60 * time.Second
	pingPeriod     = (pongWait * 9) / 10
	maxMessageSize = 512
)

var upgrader = websocket.Upgrader{}

type WSClient struct {
	conn       *websocket.Conn
	player     *data.Player
	output     chan []byte
	remoteHost string
	hub        *WSHub
}

type errorMessage struct {
	Cmd   string `json:"cmd"`
	Error string `json:"error"`
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
		_, packet, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				log.Printf("error: %v", err)
			}
			break
		}
		HandleMessage(c, packet)
	}
}

func (c *WSClient) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case packet, ok := <-c.output:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := c.conn.WriteMessage(websocket.TextMessage, packet)
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

func (c *WSClient) Error(format string, v ...interface{}) {
	c.Send(errorMessage{"error", fmt.Sprintf(format, v...)})
}

func (c *WSClient) Broadcast(v interface{}) {
	hub.Broadcast(v)
}

func (c *WSClient) Send(v interface{}) {
	g, err := json.Marshal(v)
	if err != nil {
		return
	}
	c.output <- g
}

func (c *WSClient) Host() string {
	return c.remoteHost
}

func (c *WSClient) SetPlayer(p *data.Player) {
	c.player = p
}

func (c *WSClient) GetPlayer() *data.Player {
	return c.player
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

	handleConnect(client)
}
