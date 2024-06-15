package main

import "log"

const topic = "obu"

func main() {
	logger := NewLogrusLogger()
	service := NewCalculatorService()
	service = NewLogMiddleware(logger, service)
	consumer, err := NewKafkaConsumer(topic, service)
	if err != nil {
		log.Fatal(err)
	}

	consumer.Start()
}
