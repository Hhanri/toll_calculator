package main

import "log"

const topic = "obu"

func main() {
	service := NewCalculatorService()
	consumer, err := NewKafkaConsumer(topic, service)
	if err != nil {
		log.Fatal(err)
	}

	consumer.Start()
}
