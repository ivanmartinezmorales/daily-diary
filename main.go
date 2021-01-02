package main

import (
	"log"

	"github.com/ivanmartinezmorales/life-server/server"
)

func main() {
	app := server.CreateServer()
	log.Fatal(app.Listen(":3000"))
}
