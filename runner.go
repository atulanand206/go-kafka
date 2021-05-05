package kafka

import (
	"strings"

	"github.com/namsral/flag"
)

var (
	listenAddrApi      string
	kafkaBrokerUrl     string
	kafkaVerbose       bool
	kafkaTopic         string
	kafkaConsumerGroup string
	kafkaClientId      string
)

func LoadConsumer(brokerId string, topic string) {
	flag.StringVar(&kafkaBrokerUrl, "kafka-brokers", brokerId, "Kafka brokers in comma separated value")
	flag.StringVar(&kafkaTopic, "kafka-topic", topic, "Kafka topic. Only one topic per worker.")
	flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaConsumerGroup, "kafka-consumer-group", "consumer-group", "Kafka consumer group")
	flag.StringVar(&kafkaClientId, "kafka-client-id", "kafka-client-id", "Kafka client id")

	flag.Parse()

	ConfigureReader(strings.Split(kafkaBrokerUrl, ","), kafkaClientId, kafkaTopic)
}

func LoadPublisher(brokerId string, topic string) {
	flag.StringVar(&kafkaBrokerUrl, "kafka-brokers", brokerId, "Kafka brokers in comma separated value")
	flag.StringVar(&kafkaTopic, "kafka-topic", topic, "Kafka topic to push")
	flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaClientId, "kafka-client-id", "kafka-client-id", "Kafka client id to connect")

	flag.Parse()

	ConfigureProducer(strings.Split(kafkaBrokerUrl, ","), kafkaClientId, kafkaTopic)
}
