package main

import (
	"github.com/gofiber/fiber/v3"
)

func main() {
	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("Hello, World! Auth Service")
	})

	app.Listen(":8001")
}
