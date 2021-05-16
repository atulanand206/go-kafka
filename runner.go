package kafka

import (
	"strings"

	"github.com/namsral/flag"
)

var (
	// Contains the broker url for the kafka cluster.
	KafkaBrokerUrl string
	// Used for printing the verbose logs to the topic. Set to true for now.
	kafkaVerbose bool
	// Contains the topic used for interacting with the broker.
	KafkaTopic string
	// Contains the consumer group present in the cluster. Fixed for now.
	kafkaConsumerGroup string
	// Contains the client id of the broker. Fixed for now.
	kafkaClientId string
)

/* Configures the Reader to register a subscriber for the the kafka broker and topic.*/
func LoadConsumer(brokerId string, topic string) {
	flag.StringVar(&KafkaBrokerUrl, "kafka-brokers", brokerId, "Kafka brokers in comma separated value")
	flag.StringVar(&KafkaTopic, "kafka-topic", topic, "Kafka topic. Only one topic per worker.")
	flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaConsumerGroup, "kafka-consumer-group", "consumer-group", "Kafka consumer group")
	flag.StringVar(&kafkaClientId, "kafka-client-id", "kafka-client-id", "Kafka client id")

	flag.Parse()

	ConfigureReader(strings.Split(KafkaBrokerUrl, ","), kafkaClientId, KafkaTopic)
}

/* Configures the Writer to register a publisher for the the kafka broker and topic.*/
func LoadPublisher(brokerId string, topic string) {
	flag.StringVar(&KafkaBrokerUrl, "kafka-brokers", brokerId, "Kafka brokers in comma separated value")
	flag.StringVar(&KafkaTopic, "kafka-topic", topic, "Kafka topic to push")
	flag.BoolVar(&kafkaVerbose, "kafka-verbose", true, "Kafka verbose logging")
	flag.StringVar(&kafkaClientId, "kafka-client-id", "kafka-client-id", "Kafka client id to connect")

	flag.Parse()

	ConfigureWriter(strings.Split(KafkaBrokerUrl, ","), kafkaClientId, KafkaTopic)
}
