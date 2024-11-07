package route

import (
	"meli/handlers"

	"github.com/gofiber/fiber/v2"
)

func ShowRoute(app *fiber.App) {
	app.Get("/shows", handlers.ListShows)
}
