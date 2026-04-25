package main

import (
	"encoding/json"
	"log"
	"time"

	"github.com/gofiber/fiber/v3"
	"github.com/pnaskardev/EventLens/auth-service/utils"
	internal_config "github.com/pnaskardev/EventLens/internal/config"
	"github.com/pnaskardev/EventLens/internal/constants"
	"github.com/pnaskardev/EventLens/internal/kafka"
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

	// Run a separate go routine to generate logs
	go func() {
		for {
			for _, authLog := range utils.AuthLogs {
				time.Sleep(5 * time.Second)

				authLog.Timestamp = time.Now()

				logByte, err := json.Marshal(authLog)
				if err != nil {
					log.Printf("error while marshalling log %v\n", err)
				}

				producer.Publish(constants.AUTH_SERVICE_LOGS_TOPIC, logByte)

			}
		}
	}()

	app := fiber.New()

	app.Get("/health", func(c fiber.Ctx) error {
		return c.SendString("Hello, World! Auth Service")
	})

	app.Listen(":8001")
}
