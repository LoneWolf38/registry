package main

import (
	"log"
	"net"
)

// TCP

type Registry struct {
	conn net.Conn
}

const (
	OPCODE_LEN = 2
	// stores the opcode, which will be upto 2 bytes
	// Opcode can be the
	// 1. Register(RE) - takes in the hostname, ip, mac-address, uptime, type of os, timestamp when its sending
	// 2. Status(ST) - takes in the hostname, ip, mac-address, uptime,type of os, timestamp when its sending
	// 3. DeRegister(DR) - takes in the hostname, ip, mac-address, uptime, type of os, timestamp when its sending
	// 4. Shutdown(SH) - takes in the hostname, ip, mac-address, uptime, type of os, timestamp when its sending
	//
	//
	//
	//
	// Host Address taken in a 12 byte address of the client host
	IP_ADDRESS_LEN = 12
)

func parsePackets(conn net.Conn) {
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
