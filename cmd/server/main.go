package main

import (
	"flag"
	"log"

	"github.com/m-nny/goinit/pkg/mcserver"
)

var (
	host = flag.String("host", "localhost", "ip host")
	port = flag.Uint("port", 8080, "port")
)

func main() {
	server := mcserver.NewServer()
	if err := server.Start(*host, *port); err != nil {
		log.Fatal(err)
	}
	defer server.Close()
}
