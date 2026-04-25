package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	internal_config "github.com/pnaskardev/EventLens/internal/config"
	"github.com/pnaskardev/EventLens/internal/constants"
	"github.com/pnaskardev/EventLens/internal/kafka"
	"github.com/pnaskardev/EventLens/order-service/utils"
)

func main() {

	config, err := internal_config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}

	producer, err := kafka.NewKafkaProducer(config.KafkaBrokers)
	if err != nil {
		log.Printf("couldnt intialize kafka producer: %v\n", err)
	}

	go func() {
		for {
			for _, orderLog := range utils.OrderLogs {
				time.Sleep(10 * time.Second)
				orderLog.Timestamp = time.Now()

				logByte, err := json.Marshal(orderLog)
				if err != nil {
					log.Printf("error while marshalling log %v\n", err)
				}

				producer.Publish(constants.ORDER_SERVICE_LOGS_TOPIC, logByte)
			}
		}
	}()

	app := fiber.New()
	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("Hello, World! Order Service")
	})

	app.Listen(":8002")
}
