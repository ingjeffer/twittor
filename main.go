package main

import (
	"log"

	"github.com/ingjeffer/twittor/bd"
	"github.com/ingjeffer/twittor/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la BD")
		return
	}

	handlers.Manejadores()
}
