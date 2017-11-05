package main

import (
	"flag"

	"github.com/bpa/tables"
)

var addr = flag.String("addr", ":8080", "http service address")

func main() {
	flag.Parse()
	tables.Listen(*addr)
}
