package main

import (
	"meli/database"
	"meli/handlers"
	route "meli/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := database.DbConnection()

	app := fiber.New()

	handlers.SetDB(db)

	route.ShowRoute(app)
	route.ReservationRoute(app)

	app.Listen(":3000")
}
