package main

import (
	"log"
	"src/codetwitter/handlers"

	"src/codetwitter/bd"
)

func main() {
	if bd.ChequeoConection() == 0 {
		log.Fatal("Sin conexión a la BD")
		return
	}
	handlers.Manejadores()
}
