package server

import (
	"log"
	"net"
)

func NewServer(listenAddr string) *Server {
	return &Server{
		ListenAddr: listenAddr,
		Quitch:     make(chan struct{}),
		Msgch:      make(chan Message, 10),
	}
}

func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	defer ln.Close()
	s.Ln = ln
	go s.acceptLoop()
	// wait for quite channel
	<-s.Quitch
	close(s.Msgch)

	return nil
}

func (s *Server) acceptLoop() {
	for {
		conn, err := s.Ln.Accept()
		if err != nil {
			log.Println("accept error:", err)
			continue
		}
		log.Println("new connection to server:", conn.RemoteAddr())
		go s.readLoop(conn)
	}
}

func (s *Server) readLoop(conn net.Conn) {
	defer conn.Close()
	buf := make([]byte, 2048)
	for {
		n, err := conn.Read(buf)
		if err != nil {
			log.Println("read error", err)
			continue
		}
		s.Msgch <- Message{
			From:    conn.RemoteAddr().String(),
			Payload: buf[:n],
		}
	}
}
