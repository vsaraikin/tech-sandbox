package main

import (
	"log"

	"github.com/IBM/sarama"
)

const (
	broker1 = "localhost:9092"
	topic   = "example"
)

func main() {
	// Setup configuration
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// Create a new consumer
	brokers := []string{broker1}
	consumer, err := sarama.NewConsumer(brokers, config)
	if err != nil {
		log.Fatalln("Failed to start Sarama consumer:", err)
	}
	defer consumer.Close()

	// Consume a message
	partitionConsumer, err := consumer.ConsumePartition(topic, 0, sarama.OffsetNewest)
	if err != nil {
		log.Fatalln("Failed to start Sarama partition consumer:", err)
	}
	defer partitionConsumer.Close()

	// Listen for messages
	for {
		select {
		case msg := <-partitionConsumer.Messages():
			log.Printf("Received message: %s\n", string(msg.Value))
		case err := <-partitionConsumer.Errors():
			log.Println("Failed to read message:", err)
		}
	}
}
