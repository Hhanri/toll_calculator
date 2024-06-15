package main

import (
	"encoding/json"
	"fmt"

	"github.com/Hhanri/toll_calculator/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorServicer
}

func NewKafkaConsumer(topic string, service CalculatorServicer) (*KafkaConsumer, error) {
	c, err := kafka.NewConsumer(
		&kafka.ConfigMap{
			"bootstrap.servers": "localhost",
			"group.id":          "group",
			"auto.offset.reset": "earliest",
		},
	)

	if err != nil {
		return nil, err
	}

	c.SubscribeTopics([]string{topic}, nil)

	return &KafkaConsumer{
		consumer:    c,
		isRunning:   true,
		calcService: service,
	}, nil
}

func (c *KafkaConsumer) Start() {
	fmt.Println("kafka consumer started")
	c.readMessageLoop()
}

func (c *KafkaConsumer) readMessageLoop() {
	for c.isRunning {

		msg, err := c.consumer.ReadMessage(-1)

		if err != nil {
			fmt.Printf("kafka consume error: %s\n", err)
			continue
		}

		var data types.ObuData
		if err := json.Unmarshal(msg.Value, &data); err != nil {
			fmt.Printf("JSON serialization error: %s\n", err)
			continue
		}
		_, err = c.calcService.CalculateDistance(data)
		if err != nil {
			fmt.Printf("Calculation error: %s\n", err)
			continue
		}
	}
}
