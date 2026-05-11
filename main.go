package main

import (
	"log"

	"backend-go/database"
	"backend-go/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/stripe/stripe-go/v74"
)

func main() {
	stripe.Key = "sk_test_your_stripe_secret_key_here"

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:3000",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	database.Init()
	routes.Setup(app)

	log.Println("Server running on port 8080")
	app.Listen(":8080")
}
