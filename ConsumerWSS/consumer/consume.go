package consumer

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/M1keTrike/EventDriven/models"
	"github.com/M1keTrike/EventDriven/sender"
	"github.com/gorilla/websocket"
	"github.com/rabbitmq/amqp091-go"
)

func Consume(conn *amqp091.Connection, failOnError func(error, string), wsconn *websocket.Conn) {
	cq := os.Getenv("CONSUME_QUEUE")
	wsURL := os.Getenv("WEB_SOCKET_URL")

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	msgs, err := ch.Consume(
		cq,
		"ConsumerWebSocketS",
		true,
		false,
		false,
		false,
		nil,
	)
	failOnError(err, "Failed to register a consumer")

	sender := sender.NewSender(wsURL)

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Printf("Received a message: %s", d.Body)

			var receivedMsg models.MessageRecieveRMQ
			err := json.Unmarshal(d.Body, &receivedMsg)
			if err != nil {
				log.Printf("Error decoding message: %v", err)
				continue
			}

			content, err := json.Marshal(receivedMsg)
			if err != nil {
				log.Printf("Error encoding message to JSON: %v", err)
				continue
			}

			messageWS := models.NewMessageWS(
				fmt.Sprintf("%d", receivedMsg.UserId),
				fmt.Sprintf("%d", receivedMsg.UserId),
				string(content),
				receivedMsg.Time,
			)

			messageJSON, err := json.Marshal(messageWS)
			if err != nil {
				log.Printf("Error encoding WebSocket message: %v", err)
				continue
			}

			err = sender.Send(messageJSON, wsconn)
			if err != nil {
				log.Printf("Error sending message to WebSocket: %v", err)
			}
		}
	}()

	log.Println(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever
}
