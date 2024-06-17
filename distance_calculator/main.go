package main

import (
	"log"

	"github.com/Hhanri/toll_calculator/aggregator/client"
)

const topic = "obu"
const aggEndPoint = "http://127.0.0.1:4000/aggregate"

func main() {
	client := client.NewClient(aggEndPoint)
	logger := NewLogrusLogger()
	service := NewCalculatorService()
	service = NewLogMiddleware(logger, service)
	consumer, err := NewKafkaConsumer(topic, service, client)
	if err != nil {
		log.Fatal(err)
	}

	consumer.Start()
}
