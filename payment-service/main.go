package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/pnaskardev/EventLens/payment-service/utils"
)

func main() {

	go func() {
		for {
			for _, paymentLog := range utils.PaymentLogs {
				time.Sleep(12 * time.Second)
				paymentLog.Timestamp = time.Now()
				log.Println(paymentLog)
			}
		}
	}()

	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("Hello, World! Payment Service")
	})

	app.Listen(":8003")
}
