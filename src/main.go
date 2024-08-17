package main

import (
	"log"
	"os"

	"github.com/adelylria/Game-of-life-API-Go/server"
	"github.com/joho/godotenv"
)

func main() {

	// Carga las variables de entorno desde el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Fatalf("Error al cargar el archivo .env: %v", err)
	}
	// Lee el puerto desde una variable de entorno, si está configurado
	port := os.Getenv("PORT")

	if port == "" {
		port = ":8080" // Usa el puerto por defecto si no se especifica otro
	}

	srv := server.NewServer()

	// Inicia el servidor
	log.Printf("Iniciando servidor en el puerto %s...\n", port)

	if err := srv.Run(port); err != nil {
		log.Fatalf("Error al iniciar el servidor: %v", err)
	}
}
