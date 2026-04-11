package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/pnaskardev/EventLens/auth-service/utils"
)

func main() {

	// Run a separate go routine to generate logs
	go func() {
		for {
			for _, authLog := range utils.AuthLogs {
				time.Sleep(5 * time.Second)

				authLog.Timestamp = time.Now()
				log.Println(authLog)
			}
		}
	}()

	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("Hello, World! Auth Service")
	})

	app.Listen(":8001")
}
