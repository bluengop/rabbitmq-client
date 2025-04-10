package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	rabbithole "github.com/michaelklishin/rabbit-hole/v3"
)

/*
Management API: https://rawcdn.githack.com/rabbitmq/rabbitmq-server/v4.0.7/deps/rabbitmq_management/priv/www/api/index.html
RabbitHole: 	https://pkg.go.dev/github.com/michaelklishin/rabbit-hole/v3
				https://github.com/michaelklishin/rabbit-hole
*/

func main() {
	// Load credentials
	err := godotenv.Load("secrets.env")
	if err != nil {
		panic("Error loading secrets.env file")
	}
	endpoint := os.Getenv("RMQ_ENDPOINT")
	user := os.Getenv("RMQ_USER")
	pass := os.Getenv("RMQ_PASS")

	rmqc, err := createRabbitMQClient(endpoint, user, pass)
	if err != nil {
		log.Fatalf("Unable to create RabbitMQ client: %s\n", err)
		panic(err)
	}

	queues, err := getRabbitMQQueues(rmqc)
	if err != nil {
		log.Fatalf("Unable to create RabbitMQ client: %s\n", err)
		panic(err)
	}
	for _, queue := range *queues {
		log.Printf("Queue: %s, Status: %s\n", queue.Name, queue.Status)
	}
	log.Printf("Finished.")
}

func createRabbitMQClient(endpoint, user, pass string) (*rabbithole.Client, error) {
	log.Println("Creating new RabbitMQ client")
	client, err := rabbithole.NewClient(endpoint, user, pass)
	if err != nil {
		log.Fatalf("Error creating RabbitMQ client: %s\n", err)
		return nil, err
	}
	return client, nil
}

func getRabbitMQQueues(client *rabbithole.Client) (*[]rabbithole.QueueInfo, error) {
	log.Printf("Getting RabbitMQ queues in %s\n", client.Endpoint)
	var queues []rabbithole.QueueInfo
	queues, err := client.ListQueues()
	if err != nil {
		log.Fatalf("Error when listing queues in the vhost: %s\n", err)
		return nil, err
	}
	return &queues, nil
}
