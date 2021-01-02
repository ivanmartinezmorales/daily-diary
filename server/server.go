package server

import (
	"log"

	"github.com/ivanmartinezmorales/life-server/server/storage"

	"github.com/gofiber/fiber"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func CreateServer() *fiber.App {
	storage.ConnectDB()
	app := fiber.New()

	app.Use(cors.New())

	log.Fatal(app.Listen(":3000"))
	return app
}
