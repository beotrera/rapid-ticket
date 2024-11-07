package route

import (
	"meli/handlers"

	"github.com/gofiber/fiber/v2"
)

func ReservationRoute(app *fiber.App) {
    app.Post("/reservations", handlers.BasicAuth ,handlers.CreateReservation)
}
