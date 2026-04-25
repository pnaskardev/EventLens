package kafka

import (
	"fmt"
	"sync"

	"github.com/IBM/sarama"
)

type KafkaProducer struct {
	producer sarama.AsyncProducer
}

var (
	kafkaOnce             sync.Once
	kafkaProducerInstance *KafkaProducer = nil
	errorOccured          error          = nil
)

func NewKafkaProducer(brokers []string) (*KafkaProducer, error) {

	fmt.Println(brokers)

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
			producer: producer,
		}
	})

	if errorOccured != nil {
		return nil, errorOccured
	}

	return kafkaProducerInstance, nil

}

func (p *KafkaProducer) Publish(topic string, message []byte) {

	// We could have implemented sync message and that will give us an ACK as well 
	// But the server needs to run and keep on waiting 
	msg := sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.ByteEncoder(message),
	}

	p.producer.Input() <- &msg

	// partition, offset, err := p.producer.SendMessage(message)
	// if err != nil {
	// 	log.Fatalf("Failed to send message: %v", err)
	// }

	// log.Printf("Message sent to partition %d at offset %d\n", partition, offset)

}
