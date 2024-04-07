package main

import (
	"apachekafka-project/producer/models"
	"fmt"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	"log"
	"time"
)

func main() {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "randomClient1",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
	}

	op := models.NewOrderPlace(p, "orderTopic")
	orders := []models.Order{
		{
			1,
			"Market Order",
			32.90,
		},
		{
			2,
			"Pizza Order",
			15.30,
		},
		{
			3,
			"Gas Order",
			15.30,
		},
	}
	for _, order := range orders {
		if err = op.PlaceOrder(&order); err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * 10)
	}
}
