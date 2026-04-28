package main

import (
	"fmt"
	"log"

	"github.com/IBM/sarama"
	"github.com/pnaskardev/EventLens/internal/config"
	"github.com/pnaskardev/EventLens/internal/constants"
	"github.com/pnaskardev/EventLens/internal/kafka"
)

func main() {
	config, err := config.InitConfig()
	if err != nil {
		log.Fatalf("error initialising config: %v", err)
	}

	consumer, err := kafka.NewKafkaConsumer(config.KafkaBrokers, config.KafkaLogGroupID)
	if err != nil {
		log.Printf("error while initalizing a consumer with group ID: %s - %v\n", config.KafkaLogGroupID, err)
	}
	fmt.Println("Consumption starting")

	topics := []string{constants.AUTH_SERVICE_LOGS_TOPIC, constants.ORDER_SERVICE_LOGS_TOPIC, constants.PAYMENT_SERVICE_LOGS_TOPIC}

	if err := consumer.Start(topics, func(msg *sarama.ConsumerMessage) error {

		fmt.Println(msg)

		return nil
	}); err != nil {
		log.Println("failed to consumed message: ", err)
	}

}
