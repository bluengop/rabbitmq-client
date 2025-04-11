package main

import (
	"fmt"
	"os"

	"github.com/bluengop/rabbitmq-client/internal/rabbitmq"
	"github.com/bluengop/rabbitmq-client/internal/utils"
	"github.com/joho/godotenv"
)

const secrets = "secrets.env"

func main() {
	// Set up logger and load credentials
	log := utils.NewLogger("info")

	if _, err := os.Stat(secrets); os.IsNotExist(err) {
		log.Fatal(fmt.Sprintf("Secrets file %s does not exist\n", secrets))
	}
	err := godotenv.Load(secrets)
	if err != nil {
		log.Fatal(fmt.Sprintf("Error loading %s file: %s\n", secrets, err))
	}
	endpoint := os.Getenv("RMQ_ENDPOINT")
	user := os.Getenv("RMQ_USER")
	pass := os.Getenv("RMQ_PASS")

	// New RabbitMQ client
	log.Info("Creating RabbitMQ client\n")
	log.Debug(fmt.Sprintf("Endpoint: %s, User: %s, Pass: %s\n", endpoint, user, pass))
	rmqc, err := rabbitmq.CreateRabbitMQClient(endpoint, user, pass)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to create RabbitMQ client: %s\n", err))
	}

	// Checking queries
	log.Info("Getting RabbitMQ queries\n")
	queues, err := rabbitmq.GetRabbitMQQueues(rmqc)
	if err != nil {
		log.Error(fmt.Sprintf("Unable to create RabbitMQ client: %s\n", err))
	}

	log.Debug("Printing queue info:\n")
	failed := 0
	for _, queue := range *queues {
		log.Debug(fmt.Sprintf("Queue: %s, Status: %s\n", queue.Name, queue.Status))
		if queue.Status != "running" {
			log.Warn(fmt.Sprintf("Queue not running: Name: %s, Status: %s\n", queue.Name, queue.Status))
			failed++
		}
	}
	if failed == 0 {
		log.Info("All queues are running ✅\n")
	}
	log.Info("Finished")
}
