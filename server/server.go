package server

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/ivanmartinezmorales/life-server/server/storage"
)

func handleGetStatus(c *fiber.Ctx) error {
	msg := "The system is up"
	// Do some business logic in here
	return c.SendString(msg)
}

func handleGetAllDiaryEntries(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Fetching all records")
	// Do some business logic in here
	return c.SendString(msg)
}

func handleGetDiaryEntry(c *fiber.Ctx) error {
	msg := fmt.Sprintf("Fetched a record with the id of: %s", c.Params("id"))
	// Do some business logic in here
	return c.SendString(msg)
}

func setupRoutes(app *fiber.App) {
	api := app.Group("api/v1")
	api.Get("/", handleGetStatus)
	api.Get("/entries", handleGetAllDiaryEntries)
	api.Get("/entries/:id", handleGetDiaryEntry)
}

// CreateServer builds the server
func CreateServer() *fiber.App {
	storage.ConnectDB()

	app := fiber.New()
	app.Use(cors.New())
	app.Use(logger.New(logger.Config{
		Format: "${locals:requestid} ${status} - ${method} ${path}\n",
	}))
	setupRoutes(app)

	return app
}
