# go-kafka

[![Go Reference](https://pkg.go.dev/badge/github.com/atulanand206/go-kafka.svg)](https://pkg.go.dev/github.com/atulanand206/go-kafka)

A library exposing an implementation of [segmentio/kafka-go](https://github.com/segmentio/kafka-go). You can easily get off the mark with the bare essential configurations for registering the publishers and subscribers with a Kafka cluster running.

## What is Kafka?

Kafka is a framework implementation of a software bus using stream-processing. It provides messaging queue with an interactive interface. Once you have Kafka running any service written in any language can register itself as publisher or subscriber of a topic. It can also be considered as a sequential implementation when multiple services can interact among themselves without the need of a REST endpoint.

## Pre-requisites

- You should have a kafka cluster running on your machine.
- A Kafka Topic must be registered with the cluster.

## How to implement

- [Runner.go](./runner.go) demonstrates configuring the kafka client in a way. The configurations are imported using [Flag Parser](github.com/namsral/flag) for assigning the idea to a variables. 

- You can import the package using go modules or any dependency management tool.
```go
import (
    //...
	"github.com/atulanand206/go-kafka"
    //...
)
```

- Set the topic and brokerUrl into a variable. I prefer setting the topic and broker id as environment variables.
```go
kafkaTopic := os.Getenv("KAFKA_TOPIC")
kafkaBrokerId := os.Getenv("KAFKA_BROKER_ID")
```

- Configure the Producer in your application runner method to initialize publishing to the kafka cluster. 
```go
kafka.LoadPublisher(kafkaBrokerId, kafkaTopic)
```

- Call the Push method where you'd like to publish to the kafka topic.
```go
formInBytes, _ := json.Marshal(jsonObject)
if err := kafka.Push(nil, formInBytes); err != nil {
    log.Panic(err)
}
```

- Configure the Reader in your application runner method to listen to the topic in the kafka cluster.
```go
kafka.LoadConsumer(kafkaBrokerId, kafkaTopic)
```

- Call the Read method where you'd like to listen to the kafka topic.
```go
kafka.Read(func(val string) {
    jsonObject := &object{}
    json.Unmarshal([]byte(val), &jsonObject)
}
```

## Author

- Atul Anand

