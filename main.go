package main

import (
	"flag"
	"log"
	"net/http"
)

var addr = flag.String("addr", ":8080", "http service address")
var hub Hub

func main() {
	flag.Parse()

	loadStartupFiles()
	startCleaner()

	wsHub := newWSHub()
	hub = wsHub
	go wsHub.run()

	http.HandleFunc("/websocket", func(w http.ResponseWriter, r *http.Request) {
		serveWs(wsHub, w, r)
	})
	http.Handle("/", http.FileServer(http.Dir("dist")))

	err := http.ListenAndServe(*addr, nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
