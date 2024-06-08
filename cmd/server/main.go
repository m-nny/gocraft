package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	"github.com/m-nny/goinit/pkg/mcnet"
)

var (
	host = flag.String("host", "localhost", "ip host")
	port = flag.Uint("port", 8080, "port")
)

func main() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *host, *port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	log.Printf("Server started at %s:%d", *host, *port)

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("Got connection: %v\n", conn)
		go func() {
			err := handleConnectionRequest(conn)
			if err != nil {
				log.Printf("err: %v", err)
			}
		}()

	}
}

func handleConnectionRequest(conn net.Conn) error {
	defer conn.Close()

	packet, err := mcnet.ReadGenericPacket(conn)
	if err != nil {
		return fmt.Errorf("err reading packet: %v", err)
	}

	handshake, ok := packet.(*mcnet.HandshakePacket)
	if !ok {
		return fmt.Errorf("expected Handshake Packet, got: %+v", packet)
	}
	log.Printf("got handshake packet: %+v", handshake)

	packet, err = mcnet.ReadGenericPacket(conn)
	if err != nil {
		return err
	}

	return nil
}
