package kafka

import (
	"sync"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	Producer *sarama.AsyncProducer
}

var (
	kafkaOnce             sync.Once
	kafkaProducerInstance *KafkaProducer = nil
	errorOccured          error          = nil
)

func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {

	if kafkaProducerInstance != nil {
		return kafkaProducerInstance, nil
	}

	kafkaOnce.Do(func() {
		kafkaConfig := NewKafkaProducerConfig()
		producer, err := sarama.NewAsyncProducer(brokers, kafkaConfig)
		if err != nil {
			errorOccured = err
			return
		}

		kafkaProducerInstance = &KafkaProducer{
			Producer: &producer,
		}
	})

	if errorOccured != nil {
		return nil, errorOccured
	}

	return kafkaProducerInstance, nil

}
