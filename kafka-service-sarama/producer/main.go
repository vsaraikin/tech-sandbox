package main

import (
	"github.com/IBM/sarama"

	"log"
)

const (
	broker1 = "localhost:9092"
	topic   = "example"

	exampleMessage = "Hello, World!"
)

func main() {
	// Setup configuration
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true

	// Create a new producer
	brokers := []string{broker1}
	producer, err := sarama.NewSyncProducer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama producer:", err)
	}
	defer producer.Close()

	// Construct a message
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(exampleMessage),
	}

	// Send the message
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		log.Fatalln("Failed to send message:", err)
	}
	log.Printf("Message sent to partition %d at offset %d\n", partition, offset)
}
