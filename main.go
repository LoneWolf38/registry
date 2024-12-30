package main

import (
	"log"
	"net"
)

func main() {
	cn, err := net.Listen("tcp", "0.0.0.0:6969")
	if err != nil {
		log.Fatal(err.Error())
	}
	for {
		conn, err := cn.Accept()
		if err != nil {
			log.Fatal(err.Error())
		}
		go handleConn(conn)
	}
}
