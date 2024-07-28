package main

import (
	"bytes"
	"fmt"
	"log"

	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/m-nny/goinit/pkg/datatypes"
)

func test1(input []byte) {
	r := bytes.NewReader(input)

	fmt.Println("===============")
	fmt.Println("test1")
	fmt.Println("===============")
	var (
		Protocol, Intention pk.VarInt
		ServerAddress       pk.String        // ignored
		ServerPort          pk.UnsignedShort // ignored
	)
	fmt.Printf("unread: %d\n", r.Len())
	// receive handshake packet
	var p pk.Packet
	if err := p.UnPack(r, -1); err != nil {
		log.Fatal(err)
	}
	if err := p.Scan(&Protocol, &ServerAddress, &ServerPort, &Intention); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("procotol %v\n", Protocol)
	fmt.Printf("address %v\n", ServerAddress)
	fmt.Printf("ServerPort %v\n", ServerPort)
	fmt.Printf("Intention %v\n", Intention)
	fmt.Printf("unread: %d\n", r.Len())
}

func test2(input []byte) {
	fmt.Println("===============")
	fmt.Println("test2")
	fmt.Println("===============")

	r := bytes.NewReader(input)
	var packetLen datatypes.VarInt
	n, err := packetLen.ReadFrom(r)
	fmt.Printf("n: %d err: %v\n", n, err)
	fmt.Printf("packetLen: %+v\n", packetLen)

	var packetId datatypes.VarInt
	n, err = packetId.ReadFrom(r)
	fmt.Printf("n: %d err: %v\n", n, err)
	fmt.Printf("packetId: %+v\n", packetId)
}

func main() {
	input := ([]byte{16, 0, 254, 5, 9, 108, 111, 99, 97, 108, 104, 111, 115, 116, 31, 144, 1, 1, 0})
	fmt.Printf("%+v\n", input)
	test1(input)
	test2(input)
}
