package main

import (
	"fmt"

	rabbithole "github.com/michaelklishin/rabbit-hole/v3"
)

/*
Management API: https://rawcdn.githack.com/rabbitmq/rabbitmq-server/v4.0.7/deps/rabbitmq_management/priv/www/api/index.html
RabbitHole: 	https://pkg.go.dev/github.com/michaelklishin/rabbit-hole/v3
				https://github.com/michaelklishin/rabbit-hole
*/

func CreateRabbitMQClient(endpoint, user, pass string) *rabbithole.Client {
	fmt.Println("Creating RabbitMQ client")
	client, err := rabbithole.NewClient(endpoint, user, pass)
	if err != nil {
		fmt.Printf("Error creating RabbitMQ client: %s\n", err)
		panic(err)
	}
	return client
}

func GetRabbitMQqueues(client *rabbithole.Client, queueName string) {
	fmt.Println("Getting RabbitMQ queues")
	queues, err := client.GetQueues()
	if err != nil {
		panic(err)
	}
	for _, queue := range queues {
		fmt.Printf("Queue: %s, State: %d\n", queue.Name, queue.State)
	}
}
