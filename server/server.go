package server

import (
	"log"
	"net"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(addr string) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return err
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println("Error al aceptar la conexión: ", err)
			continue
		}
		go s.handleNewClient(conn)
	}
}

func (s *Server) handleNewClient(conn net.Conn) {
	client := NewClient(conn)
	defer client.Close()

	log.Printf("Nuevo cliente conectado desde %s\n", conn.RemoteAddr().String())

	welcomomeMsg := "Bienvennido al servidor TCP!\n"
	client.conn.Write([]byte(welcomomeMsg))
	client.HandleMessages()
}
