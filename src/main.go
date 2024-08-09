package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/adelylria/Game-of-life-API-Go/handlers"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	http.HandleFunc("/api/game", handlers.GameHandler)
	http.ListenAndServe(":8080", nil)
}
