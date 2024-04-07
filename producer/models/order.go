package models

import (
	"encoding/json"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
)

type Order struct {
	OrderId int32   `json:"order_id"`
	Name    string  `json:"name"`
	Price   float64 `json:"price"`
}

type OrderPlace struct {
	producer   *kafka.Producer
	topic      string
	deliveryCh chan kafka.Event
}

func NewOrderPlace(p *kafka.Producer, topic string) *OrderPlace {
	return &OrderPlace{
		producer:   p,
		topic:      topic,
		deliveryCh: make(chan kafka.Event, 10000),
	}
}

func (op *OrderPlace) PlaceOrder(order *Order) error {
	formatPayload, err := json.Marshal(order)
	if err != nil {
		log.Fatal(err)
	}

	err = op.producer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &op.topic, Partition: kafka.PartitionAny},
		Value:          formatPayload,
	},
		op.deliveryCh,
	)
	if err != nil {
		log.Fatal(err)
	}
	<-op.deliveryCh
	fmt.Printf("placed order on the queue %s\n", formatPayload)

	return nil
}
