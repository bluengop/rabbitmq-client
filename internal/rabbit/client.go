package rabbit

import (
	"log"

	rabbithole "github.com/michaelklishin/rabbit-hole/v3"
)

func CreateRabbitMQClient(endpoint, user, pass string) (*rabbithole.Client, error) {
	log.Println("Creating new RabbitMQ client")
	client, err := rabbithole.NewClient(endpoint, user, pass)
	if err != nil {
		log.Fatalf("Error creating RabbitMQ client: %s\n", err)
		return nil, err
	}
	return client, nil
}

func GetRabbitMQQueues(client *rabbithole.Client) (*[]rabbithole.QueueInfo, error) {
	log.Printf("Getting RabbitMQ queues in %s\n", client.Endpoint)
	var queues []rabbithole.QueueInfo
	queues, err := client.ListQueues()
	if err != nil {
		log.Fatalf("Error when listing queues in the vhost: %s\n", err)
		return nil, err
	}
	return &queues, nil
}
