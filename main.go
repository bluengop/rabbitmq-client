package main

import (
	"fmt"
	"os"

	"github.com/bluengop/rabbitmq-client/internal/rabbit"
	"github.com/bluengop/rabbitmq-client/internal/utils"
	"github.com/joho/godotenv"
)

const secrets = "secrets.env"

func main() {
	// Set up logger and load credentials
	log := utils.CreateLogger()

	if _, err := os.Stat(secrets); os.IsNotExist(err) {
		log.Fatal(fmt.Sprintf("Secrets file %s does not exist\n", secrets))
		//panic(err)
	}
	err := godotenv.Load(secrets)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error loading %s file: %s\n", secrets, err))
	}
	endpoint := os.Getenv("RMQ_ENDPOINT")
	user := os.Getenv("RMQ_USER")
	pass := os.Getenv("RMQ_PASS")

	rmqc, err := rabbit.CreateRabbitMQClient(endpoint, user, pass)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to create RabbitMQ client: %s\n", err))
	}

	queues, err := rabbit.GetRabbitMQQueues(rmqc)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to create RabbitMQ client: %s\n", err))
	}
	for _, queue := range *queues {
		log.Debug("Printing queue info:\n")
		log.Debug(fmt.Sprintf("Queue: %s, Status: %s\n", queue.Name, queue.Status))
		if queue.Status != "running" {
			log.Info(fmt.Sprintf("Queue not running: Name: %s, Status: %s\n", queue.Name, queue.Status))
		}
	}
	log.Info("Finished")
}
