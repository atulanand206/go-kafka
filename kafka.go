package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

// Instance variable Writer to be used for writing to a kafka topic.
// Must be configured before calling the Push method.
var Writer *kafka.Writer

// Instance variable Reader to be used for reading from a kafka topic.
// Must be configured before calling the Read method.
var Reader *kafka.Reader

// Configures the Kafka reader for a given topic listening to the broker url.
func ConfigureReader(kafkaBrokerUrls []string, clientId string, topic string) {
	config := kafka.ReaderConfig{
		Brokers:         kafkaBrokerUrls,
		GroupID:         clientId,
		Topic:           topic,
		MinBytes:        10e3,            // 10KB
		MaxBytes:        10e6,            // 10MB
		MaxWait:         1 * time.Second, // Maximum amount of time to wait for new data to come when fetching batches of messages from kafka.
		ReadLagInterval: -1,
	}
	Reader = kafka.NewReader(config)
}

// Listens to the configured Kafka topic and acts every time there is a message in the partition.
func Read(handler func(value string)) {
	if Reader == nil {
		return
	}
	for {
		m, err := Reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error while receiving message: %s", err.Error())
			continue
		}

		value := string(m.Value)
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, value)
		handler(value)
	}
}

// Configures the Kafka writer for a given topic publishing to the broker url.
func ConfigureWriter(kafkaBrokerUrls []string, clientId string, topic string) {
	dialer := &kafka.Dialer{
		Timeout:  10 * time.Second,
		ClientID: clientId,
	}
	config := kafka.WriterConfig{
		Brokers:          kafkaBrokerUrls,
		Topic:            topic,
		Balancer:         &kafka.LeastBytes{},
		Dialer:           dialer,
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	Writer = kafka.NewWriter(config)
}

// Pushes to the configured Kafka topic as a message in the partition.
func Push(key, value []byte) (err error) {
	if Writer == nil {
		return
	}

	parent := context.Background()
	defer parent.Done()

	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}

	return Writer.WriteMessages(parent, message)
}
