package kafka

import (
	"context"
	"log"

	"github.com/IBM/sarama"
)

// This is custom struct which has my custome handler
// I am able to pass it in the consume method cause it implements all of the methods
// which ConsumerGroupHandler needs
type kafkaConsumerMessageHandler func(msg *sarama.ConsumerMessage) error

type kafkaGroupHandler struct {
	handler kafkaConsumerMessageHandler
}

// These are lifecycle methods.
// Currently returning nil as there are no errors

func (h *kafkaGroupHandler) Setup(sarama.ConsumerGroupSession) error {
	return nil
}

func (h *kafkaGroupHandler) Cleanup(sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim is called internally by sarama itself
func (h *kafkaGroupHandler) ConsumeClaim(
	session sarama.ConsumerGroupSession,
	claim sarama.ConsumerGroupClaim,
) error {
	for msg := range claim.Messages() {

		if err := h.handler(msg); err != nil {
			log.Println("Message handling failed:", err)
			continue
		}

		session.MarkMessage(msg, "")
	}

	return nil
}

// CONSUMER INSTANCE
type KafkaConsumer struct {
	consumerGroup sarama.ConsumerGroup
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
			consumerGroup: consumer,
		}
	})

	if errorOccured != nil {
		return nil, errorOccured
	}

	return kafkaConsumerInstance, nil

}

func (c *KafkaConsumer) Start(topics []string, handler kafkaConsumerMessageHandler) error {

	handlerMethod := &kafkaGroupHandler{
		handler: handler,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c.consumerGroup.Consume(ctx, topics, handlerMethod)

	return nil
}
