package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Env struct {
	KafkaBrokers []string
}

func InitConfig() (Env, error) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	kafkaBrokers := os.Getenv("KAFKA_BROKERS")
	if len(kafkaBrokers) == 0 {
		return Env{}, fmt.Errorf("KAFKA_BROKERS NOT FOUND IN ENV")
	}

	env := Env{
		KafkaBrokers: []string{
			kafkaBrokers,
		},
	}

	return env, nil

}
