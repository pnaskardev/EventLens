package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

func NewKafkaProducerConfig() *sarama.Config {
	config := sarama.NewConfig()

	// Basic producer settings
	config.Producer.Return.Successes = true
	config.Producer.Return.Errors = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5

	// PRODUCER BATCHING SETTINGS
	config.Producer.Flush.Messages = 1024
	config.Producer.Flush.Frequency = 10
	config.Producer.Flush.Bytes = 5 * 1024 * 1024
	config.Producer.Flush.MaxMessages = 0

	return config
}

func NewDurableKafkaConsumerConfig() *sarama.Config {
	config := sarama.NewConfig()
	// Consumer Settings
	config.Consumer.Offsets.Initial = sarama.OffsetOldest
	config.Consumer.Offsets.AutoCommit.Enable = false
	config.Consumer.Group.Session.Timeout = 30 * time.Second
	config.Consumer.Group.Heartbeat.Interval = 10 * time.Second
	config.Consumer.Group.Rebalance.Timeout = 60 * time.Second
	config.Consumer.Fetch.Min = 1
	config.Consumer.MaxWaitTime = 500 * time.Millisecond
	config.Consumer.MaxProcessingTime = 1 * time.Second
	return config
}
