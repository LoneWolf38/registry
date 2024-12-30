package main

import (
	"fmt"
	"log"
	"net"
)

// RegEventServer is the struct which stores the connections
type RegEventServer struct {
	listner  net.Listener
	conn     net.Conn
	protocol string
	port     uint16
}

func NewEventServer(proto string, port uint16) (RegEventServer, error) {
	r := RegEventServer{}
	l, err := net.Listen(proto, fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return r, err
	}
	r.listner = l
	r.port = port
	r.protocol = proto
	return r, nil
}

func (r *RegEventServer) Start() {
	for {
		_, err := r.listner.Accept()
		if err != nil {
			log.Printf("E! cannot accept connection, Reason: %s", err.Error())
		}
	}
}

func handleConn(conn net.Conn) {
	defer conn.Close()
	data := make([]byte, 16)
	for {
		n, err := conn.Read(data)
		if err != nil {
			conn.Close()
			log.Fatal(err.Error())
		}
		log.Printf("read %d bytes", n)
	}
}
