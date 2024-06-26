package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/Hhanri/toll_calculator/aggregator/client"
	"github.com/Hhanri/toll_calculator/types"
	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type KafkaConsumer struct {
	consumer    *kafka.Consumer
	isRunning   bool
	calcService CalculatorServicer
	aggClient   *client.Client
}

func NewKafkaConsumer(topic string, service CalculatorServicer, aggClient *client.Client) (*KafkaConsumer, error) {
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
		aggClient:   aggClient,
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
		distance, err := c.calcService.CalculateDistance(data)
		if err != nil {
			fmt.Printf("Calculation error: %s\n", err)
			continue
		}

		req := types.Distance{
			Value: distance,
			ObuId: data.ObuId,
			Unix:  time.Now().Unix(),
		}
		if err := c.aggClient.AggregatetInvoice(req); err != nil {
			fmt.Printf("Aggregation error: %s\n", err)
			continue
		}
	}
}
