package main

import (
	"log"
)

const kafkaTopic = "obudata"

// type DistanceCalaculator struct {
// 	consumer DataConsumer
// }

func main() {
	svc := NewCalculatorService()

	kafkaConsumer, err := NewKafkaConsumer(kafkaTopic, svc)

	if err != nil {
		log.Fatal(err)
	}

	kafkaConsumer.Start()
}
