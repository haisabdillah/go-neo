package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/haisabdillah/golang-auth/core/config"
	"github.com/haisabdillah/golang-auth/core/delivery/handlers"
	"github.com/haisabdillah/golang-auth/core/delivery/http"
	"github.com/haisabdillah/golang-auth/core/delivery/middleware"
	"github.com/haisabdillah/golang-auth/core/infrastructure/db"
	"github.com/haisabdillah/golang-auth/core/services"
	"github.com/haisabdillah/golang-auth/pkg/logging"
	"github.com/haisabdillah/golang-auth/pkg/rabbitmq"
)

func main() {
	logging.InitLog()
	cfg := config.LoadConfig()

	//Setup queue
	if cfg.Queue.Connection != "" {
		if cfg.Queue.Connection == "rabbitmq" {
			rabbitmq.InitRabbitMQ(cfg.RabbitMQ.Username, cfg.RabbitMQ.Password, cfg.RabbitMQ.Host, cfg.RabbitMQ.Port, cfg.RabbitMQ.Vhost, cfg.RabbitMQ.Pool)
			defer rabbitmq.CloseRabbitMQ()
		}
	}

	mysql, err := db.InitMySql()
	if err != nil {
		log.Fatalf("Error init database: %s", err)
		return
	}

	services := services.NewService(mysql)
	handlers := handlers.NewHandler(services)

	r := gin.Default()
	r.Use(middleware.LogMiddleware())
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, cfg.App.Name)
	})

	http.SetupRoutes(r, handlers)
	r.Run(fmt.Sprintf(":%s", cfg.App.Port)) // listen and serve on 0.0.0.0:8080
}
