package main

import "log"

const topic = "obu"

func main() {
	consumer, err := NewKafkaConsumer(topic)
	if err != nil {
		log.Fatal(err)
	}

	consumer.Start()
}
