package server

import (
	"bufio"
	"log"
	"net"
)

type Client struct {
	conn net.Conn
}

// NewClient crea una nueva instancia de Client.
func NewClient(conn net.Conn) *Client {
	return &Client{
		conn: conn,
	}
}

func (c *Client) Close() {
	if c.conn != nil {
		c.conn.Close()
	}
}

// HandleMessages maneja la comunicación con el cliente
func (c *Client) HandleMessages() {
	scanner := bufio.NewScanner(c.conn)
	for scanner.Scan() {
		message := scanner.Text()
		log.Printf("Mensaje recibido del cliente: %s\n", message)

		if message == "quit" || message == "Quit" {
			log.Println("El cliente solicitó finalizar la conexión.")
			c.conn.Write([]byte("Conexión finalizada. Adiós!\n"))
			break
		}

		// Responder al cliente
		response := "Mensaje recibido: " + message + "\n"
		c.conn.Write([]byte(response))
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error al leer del cliente: %v\n", err)
	}

	// Cierra la conexión con el cliente
	c.Close()
	log.Println("Conexión con el cliente cerrada.")
}
