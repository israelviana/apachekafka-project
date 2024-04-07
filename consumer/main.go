package main

import (
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

func main() {
	topic := "orderTopic"
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "randomGroup1",
		"auto.offset.reset": "smallest"})
	if err != nil {
		log.Fatal(err)
	}

	err = consumer.Subscribe(topic, nil)
	if err != nil {
		log.Fatal(err)
	}

	for {
		ev := consumer.Poll(100)
		switch e := ev.(type) {
		case *kafka.Message:
			//received one order
			fmt.Printf("%s\n", string(e.Value))
		case *kafka.Error:
			fmt.Printf("%v\n", e)
		}
	}
}
