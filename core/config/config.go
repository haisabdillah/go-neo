package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Config holds the application configuration
type Config struct {
	App      AppConfig
	Database DatabaseConfig
	Logging  LoggingConfig
	RabbitMQ RabbitMQ
	Queue    QueueConfig
}

// AppConfig holds App-related settings
type AppConfig struct {
	Port string
	Name string
}

type RabbitMQ struct {
	Username string
	Password string
	Host     string
	Port     string
	Vhost    string
	Pool     string
}

// DatabaseConfig holds database-related settings
type DatabaseConfig struct {
	Host          string
	Port          string
	User          string
	Password      string
	DBName        string
	DBPoolMax     string
	DBPoolIddle   string
	DBPoolMaxTime string
}

// LoggingConfig holds logging-related settings
type LoggingConfig struct {
	Level string
}

type QueueConfig struct {
	Connection string
}

// LoadConfig loads the configuration from environment variables
func LoadConfig() *Config {
	// Load .env file
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}

	return &Config{
		App: AppConfig{
			Port: getEnv("APP_PORT", "8080"),   // Default to 8080 if not set
			Name: getEnv("APP_NAME", "Golang"), // Default to 8080 if not set
		},
		Database: DatabaseConfig{
			Host:          getEnv("DB_HOST", "localhost"),
			Port:          getEnv("DB_PORT", "3306"),
			User:          getEnv("DB_USER", "root"),
			Password:      getEnv("DB_PASSWORD", "root"),
			DBName:        getEnv("DB_NAME", "golang"),
			DBPoolMax:     getEnv("DB_POOL_MAX", "5"),
			DBPoolIddle:   getEnv("DB_POOL_IDLE", "2"),
			DBPoolMaxTime: getEnv("DB_POOL_MAX_TIME", "5"),
		},
		RabbitMQ: RabbitMQ{
			Host:     getEnv("RABBITMQ_HOST", "localhost"),
			Port:     getEnv("RABBITMQ_PORT", "5672"),
			Username: getEnv("RABBITMQ_USER", "guest"),
			Password: getEnv("RABBITMQ_PASSWORD", "guest"),
			Vhost:    getEnv("RABBITMQ_VHOST", "/"),
			Pool:     getEnv("Pool", "3"),
		},
		Logging: LoggingConfig{
			Level: getEnv("LOG_LEVEL", "info"),
		},
		Queue: QueueConfig{
			Connection: getEnv("QUEUE_CONNECTION", ""),
		},
	}
}

// getEnv retrieves the value of the environment variable or returns a default value
func getEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
