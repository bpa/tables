package server

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/bpa/tables/data"
)

var hub Hub

func Listen(port int) {
	wsHub := newWSHub()
	hub = wsHub
	go wsHub.run()

	ticker := time.NewTicker(time.Minute * 15)
	go func() {
		for range ticker.C {
			if data.DeleteExpiredTables() {
				hub.Broadcast(data.GetTables())
			}
		}
	}()

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		serveWs(wsHub, w, r)
	})
	http.Handle("/", http.FileServer(http.Dir("dist")))

	l, err := net.Listen("tcp4", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
	http.Serve(l, nil)
}
