package main

import (
	"flag"
	"log"

	"github.com/m-nny/goinit/pkg/mcnet"
)

var (
	host = flag.String("host", "localhost", "ip host")
	port = flag.Uint("port", 8080, "port")
)

func main() {
	//for _, x := range []int32{1, 2, 3, -1, -2, -3} {
	//	fmt.Printf("x: %d  x >> 1: %d x >> 2: %d\n", x, x>>1, x>>2)
	//	fmt.Printf("x: %b  x >> 1: %b x >> 2: %b\n", x, x>>1, x>>2)
	//	fmt.Println()
	//}

	server := mcnet.NewServer()
	if err := server.Start(*host, *port); err != nil {
		log.Fatal(err)
	}
	defer server.Close()
}
