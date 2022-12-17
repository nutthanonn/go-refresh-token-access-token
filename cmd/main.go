package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/nutthanonn/pkg/handlers"
	"github.com/nutthanonn/pkg/middlewares"
)

func main() {
	app := fiber.New()
	handlers := handlers.NewHandlers()

	app.Post("/login", handlers.LoginRouter)
	app.Post("/refresh", handlers.RefreshToken)

	app.Use(
		middlewares.AuthorizationRequired(),
	)

	app.Get("/profile", handlers.ProfileRouter)

	app.Listen(":3000")
}
