package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/haisabdillah/golang-auth/core/config"
	"github.com/haisabdillah/golang-auth/core/job"
	"github.com/haisabdillah/golang-auth/pkg/rabbitmq"
)

func main() {

	cfg := config.LoadConfig()
	//Setup queue
	if cfg.Queue.Connection != "" {
		if cfg.Queue.Connection == "rabbitmq" {
			rabbitmq.InitRabbitMQ(cfg.RabbitMQ.Username, cfg.RabbitMQ.Password, cfg.RabbitMQ.Host, cfg.RabbitMQ.Port, cfg.RabbitMQ.Vhost, cfg.RabbitMQ.Pool)
			defer rabbitmq.CloseRabbitMQ()
		}
	}

	// Start consuming logs
	go job.LogConsume()

	// Handle graceful shutdown
	stopChan := make(chan os.Signal, 1)
	signal.Notify(stopChan, syscall.SIGINT, syscall.SIGTERM)
	<-stopChan

	log.Println("Shutting down gracefully...")
}

//Test
