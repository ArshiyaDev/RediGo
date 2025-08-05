package main

import (
	"RediGo/internal/server"
	"log"
)

func main() {

	server := server.NewServer(":6379")

	go func() {
		for msg := range server.Msgch {
			log.Printf("received message from connection(%s):%s\n:", msg.From, string(msg.Payload))
		}
	}()
	log.Fatal(server.Start())

}
