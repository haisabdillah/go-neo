package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
	"github.com/streadway/amqp"
)

var (
	rabbitConn  *amqp.Connection
	connMutex   sync.Mutex
	channelPool []*amqp.Channel
)

func InitRabbitMQ(username string, password string, host string, port string, vhost string, poolSize string) {
	var err error
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using default environment variables")
	}
	connStr := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", username, password, host, port, vhost)

	rabbitConn, err = amqp.Dial(connStr)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}

	pool, err := strconv.Atoi(poolSize)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	// Initialize the channel pool
	for i := 0; i < pool; i++ {
		channel, err := rabbitConn.Channel()
		if err != nil {
			log.Fatalf("Failed to open a channel: %v", err)
		}
		channelPool = append(channelPool, channel)
	}
}

func getChannel() (*amqp.Channel, error) {
	connMutex.Lock()
	defer connMutex.Unlock()

	if len(channelPool) == 0 {
		return nil, fmt.Errorf("no available channels")
	}

	channel := channelPool[0]
	channelPool = channelPool[1:]
	return channel, nil
}

func returnChannel(channel *amqp.Channel) {
	connMutex.Lock()
	defer connMutex.Unlock()

	channelPool = append(channelPool, channel)
}

func CloseRabbitMQ() {
	for _, channel := range channelPool {
		channel.Close()
	}
	rabbitConn.Close()
}

type RabbitMqProps struct {
	conn *amqp.Connection
	ch   *amqp.Channel
}

func New() (*RabbitMqProps, error) {
	strDial := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", os.Getenv("RABBITMQ_USERNAME"), os.Getenv("RABBITMQ_PASSWORD"), os.Getenv("RABBITMQ_HOST"), os.Getenv("RABBITMQ_PORT"), os.Getenv("RABBITMQ_VHOST"))
	conn, err := amqp.Dial(strDial)
	if err != nil {
		defer conn.Close()
		return nil, err
	}

	ch, err := conn.Channel()
	if err != nil {
		defer ch.Close()
		return nil, err
	}
	return &RabbitMqProps{
		conn: conn,
		ch:   ch,
	}, nil
}

func SendMessage(exchange, routingKey string, b interface{}) error {
	channel, err := getChannel()
	if err != nil {
		return err
	}
	defer returnChannel(channel)

	if err := channel.ExchangeDeclare(exchange, "direct", true, false, false, false, nil); err != nil {
		return err
	}
	// Mendeklarasikan quorum queue
	args := amqp.Table{
		"x-queue-type": "quorum",
	}
	// Declare a queue
	queue, err := channel.QueueDeclare(
		routingKey, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		args,       // arguments
	)
	if err != nil {
		return err
	}
	// Bind the queue to the exchange with a routing key
	err = channel.QueueBind(
		queue.Name, // queue name
		routingKey, // routing key
		exchange,   // exchange
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		return err
	}
	body, err := json.Marshal(b)
	if err != nil {
		return err
	}
	err = channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
	if err != nil {
		return err
	}
	return nil
}

func Consume(queue string, autoAck bool) (<-chan amqp.Delivery, error) {
	channel, err := getChannel()
	if err != nil {
		return nil, err
	}
	defer returnChannel(channel)
	// Mendeklarasikan quorum queue
	args := amqp.Table{
		"x-queue-type": "quorum",
	}
	// Declare a queue
	q, err := channel.QueueDeclare(
		queue, // name
		true,  // durable
		false, // delete when unused
		false, // exclusive
		false, // no-wait
		args,  // arguments
	)
	if err != nil {
		return nil, err
	}

	// Set up a consumer
	msgs, err := channel.Consume(
		q.Name,  // queue
		"",      // consumer
		autoAck, // auto-ack
		false,   // exclusive
		false,   // no-local
		false,   // no-wait
		nil,     // args
	)
	if err != nil {
		return nil, err
	}
	return msgs, nil
}
