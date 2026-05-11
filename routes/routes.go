package routes

import (
	"backend-go/handlers"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	api := app.Group("/api")

	rooms := api.Group("/rooms")
	rooms.Get("/", handlers.GetRooms)
	rooms.Get("/:id", handlers.GetRoomByID)

	bookings := api.Group("/bookings")
	bookings.Get("/", handlers.GetBookings)
	bookings.Post("/checkout", handlers.CreateCheckoutSession)

	api.Get("/success", handlers.SuccessPayment)
}
