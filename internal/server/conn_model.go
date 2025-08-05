package server

import "net"

type Message struct {
	From    string
	Payload []byte
}

type Server struct {
	ListenAddr string
	Ln         net.Listener
	Quitch     chan struct{}
	Msgch      chan Message
}
