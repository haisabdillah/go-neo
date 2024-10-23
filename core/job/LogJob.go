package job

import (
	"bytes"
	"log"
	"net/http"

	"github.com/haisabdillah/golang-auth/pkg/rabbitmq"
)

func LogConsume() {
	msgs, err := rabbitmq.Consume("log.endpoints", true)
	if err != nil {
		log.Fatalf("error on consume %s", err)
		return
	}
	// Create a channel to signal when to stop

	// Goroutine to handle incoming messages
	go func() {
		for d := range msgs {
			log.Printf("Queue Start Endpoints Log ")
			url := "http://localhost:9200/endpoints/_doc"
			// Process the message

			req, err := http.NewRequest("POST", url, bytes.NewBuffer(d.Body))
			if err != nil {
				log.Printf("Error creating POST request: %v", err)
				return
			}
			req.Header.Set("Content-Type", "application/json")

			req.SetBasicAuth("elastic", "elastic")
			client := &http.Client{}
			resp, err := client.Do(req)
			if err != nil {
				log.Printf("Error making POST request: %v", err)
			}
			resp.Body.Close()
			log.Printf("Queue Finish Endpoints Log")
		}
	}()
}
