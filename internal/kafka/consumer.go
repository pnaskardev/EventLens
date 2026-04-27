package kafka

import (
	"context"

	"github.com/IBM/sarama"
)

type KafkaConsumer struct {
	consumerGroup *sarama.ConsumerGroup
}

var (
	kafkaConsumerInstance *KafkaConsumer = nil
)

func NewKafkaConsumer(brokers []string, kafkaGroupID string) (*KafkaConsumer, error) {

	if kafkaConsumerInstance != nil {
		return kafkaConsumerInstance, nil
	}

	kafkaOnce.Do(func() {
		kafkaConsumerConfig := NewDurableKafkaConsumerConfig()
		consumer, err := sarama.NewConsumerGroup(brokers, kafkaGroupID, kafkaConsumerConfig)
		if err != nil {
			errorOccured = err
			return
		}

		kafkaConsumerInstance = &KafkaConsumer{
			consumerGroup: &consumer,
		}
	})

	if errorOccured != nil {
		return nil, errorOccured
	}

	return kafkaConsumerInstance, nil

}

func (c *KafkaConsumer) Start(topics []string) error {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c.consumerGroup.

	return nil
}
