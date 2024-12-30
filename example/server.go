package main

import (
	"log"

	"github.com/LoneWolf38/registry/pkg/protocol"
)

/*
* This server example simulates a list of hosts which have the client running with constant updates being sent to the server for the registry storage
*
 */

const (
	TCP_PROTOCOL = "tcp"
	SERVER_PORT  = 6969
)

var hostList = []string{"192.168.1.2", "192.168.2.2", "192.168.22.23", "10.10.10.10", "10.10.10.11", "10.10.10.12"}

func main() {
	regEventList := []protocol.RegEvent{}
	deregEventlist := []protocol.RegEvent{}
	shutdownEventList := []protocol.RegEvent{}
	statusEventList := []protocol.RegEvent{}
	for _, host := range hostList {
		regEventList = append(regEventList, protocol.NewRegEvent(uint8(1), []byte("01"), []byte(host)))
		deregEventlist = append(deregEventlist, protocol.NewRegEvent(uint8(1), []byte("10"), []byte(host)))
		shutdownEventList = append(shutdownEventList, protocol.NewRegEvent(uint8(1), []byte("00"), []byte(host)))
		statusEventList = append(statusEventList, protocol.NewRegEvent(uint8(1), []byte("11"), []byte(host)))
	}
	go run_server()
	// spawn 100 clients
	log.Println("starting the server")
	run_client(regEventList, 1)
	run_client(statusEventList, 100)
	run_client(deregEventlist, 1)
	run_client(shutdownEventList, 1)
}

func run_client(events []protocol.RegEvent, freq int) {
	for _, event := range events {
		for range freq {
			c, err := protocol.NewRegEventClient(TCP_PROTOCOL, uint16(SERVER_PORT))
			if err != nil {
				log.Fatal(err)
			}
			defer c.Close()
			err = c.Send(event)
			if err != nil {
				log.Printf("cannot send data, Reason: %s", err.Error())
			}
		}
	}
}

func run_server() {
	// Create the server
	//
	s, err := protocol.NewEventServer(TCP_PROTOCOL, uint16(SERVER_PORT))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("server started")
	s.Start()
}
