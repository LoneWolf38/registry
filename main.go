package main

import (
	"log"

	"github.com/LoneWolf38/registry/pkg/protocol"
)

func main() {
	regEventServer, err := protocol.NewEventServer("tcp", 6969)
	if err != nil {
		log.Fatal(err)
	}
	regEventServer.Start()
}
