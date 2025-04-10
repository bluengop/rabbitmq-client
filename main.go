package main

import (
	"log"
	"os"

	"github.com/bluengop/rabbitmq-client/internal/rabbit"
	"github.com/joho/godotenv"
)

const secrets = "secrets.env"

func main() {
	// Load credentials
	err := godotenv.Load(secrets)
	if err != nil {
		log.Fatalf("Error loading %s file: %s\n", secrets, err)
		panic(err)
	}
	endpoint := os.Getenv("RMQ_ENDPOINT")
	user := os.Getenv("RMQ_USER")
	pass := os.Getenv("RMQ_PASS")

	rmqc, err := rabbit.CreateRabbitMQClient(endpoint, user, pass)
	if err != nil {
		log.Fatalf("Unable to create RabbitMQ client: %s\n", err)
		panic(err)
	}

	queues, err := rabbit.GetRabbitMQQueues(rmqc)
	if err != nil {
		log.Fatalf("Unable to create RabbitMQ client: %s\n", err)
		panic(err)
	}
	for _, queue := range *queues {
		log.Printf("Queue: %s, Status: %s\n", queue.Name, queue.Status)
	}
	log.Printf("Finished.")
}
