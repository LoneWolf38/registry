package main

import "log"

func main() {
	regEventServer, err := NewEventServer("tcp", 6969)
	if err != nil {
		log.Fatal(err)
	}
	regEventServer.Start()
}
