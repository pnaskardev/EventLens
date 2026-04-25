package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

type Env struct {
	KafkaBrokers []string
}

func InitConfig() (Env, error) {

	dir, err := os.Getwd()
	if err != nil {
		return Env{}, err
	}
	fmt.Printf("Current working directory: %s\n", dir)

	envPath := filepath.Join(dir, "..", ".env")

	fmt.Printf("Loading .env from: %s\n", envPath)

	if err := godotenv.Load(envPath); err != nil {
		return Env{}, fmt.Errorf("error while loading env file %s: %v", envPath, err)
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
