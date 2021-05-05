package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

var writer *kafka.Writer
var reader *kafka.Reader

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
	reader = kafka.NewReader(config)
}

func Read(handler func(value string)) {
	if reader == nil {
		return
	}
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Printf("error while receiving message: %s", err.Error())
			continue
		}

		value := string(m.Value)
		fmt.Printf("message at topic/partition/offset %v/%v/%v: %s\n", m.Topic, m.Partition, m.Offset, value)
		handler(value)
	}
}

func ConfigureProducer(kafkaBrokerUrls []string, clientId string, topic string) {
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
	writer = kafka.NewWriter(config)
}

func Push(key, value []byte) (err error) {
	if writer == nil {
		return
	}

	parent := context.Background()
	defer parent.Done()

	message := kafka.Message{
		Key:   key,
		Value: value,
		Time:  time.Now(),
	}

	return writer.WriteMessages(parent, message)
}
