package router

import (
	"github.com/gofiber/fiber/v2"
)

var users fiber.Router

// SetupRoutes is our route tree
func SetupRoutes(app *fiber.App) {
	api := app.Group("api/v1/")
	users := app.Group("/auth/v1")
}
