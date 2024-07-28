package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"

	"github.com/m-nny/goinit/pkg/packets"
)

var (
	host = flag.String("host", "localhost", "ip host")
	port = flag.Uint("port", 8080, "port")
)

func dumpAllBytes(conn net.Conn) {
	for {
		raw_packet, err := io.ReadAll(conn)
		if err != nil {
			log.Fatal(err)
		}
		if len(raw_packet) > 0 {
			log.Printf("raw_packet: %+v", raw_packet)
		}
	}
}

func dumpAllPackets(conn net.Conn) {
	for {
		p := &packets.Packet{}
		if err := p.Unpack(conn); err != nil {
			log.Fatal(err)
		}
		log.Printf("Packet: %+v", p)
	}
}

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Server started at %s:%d", *host, *port)

	for {
		rw, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}
		// go dumpAllPackets(rw)
		go dumpAllBytes(rw)
	}
}
