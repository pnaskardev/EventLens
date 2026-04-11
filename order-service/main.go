package main

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/pnaskardev/EventLens/order-service/utils"
)

func main() {

	go func() {
		for {
			for _, orderLog := range utils.OrderLogs {
				time.Sleep(10 * time.Second)
				orderLog.Timestamp = time.Now()
				log.Println(orderLog)
			}
		}
	}()

	app := fiber.New()
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("Hello, World! Order Service")
	})

	app.Listen(":8002")
}
