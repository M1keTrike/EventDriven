package main

import (
	"log"
	"os"

	"github.com/M1keTrike/EventDriven/consumer"
	"github.com/gorilla/websocket"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Panicf("%s: %s", msg, err)
	}
}

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found, using system environment variables")
	}

	rbbtURL := os.Getenv("RABBITMQ_URL")
	wsURL := os.Getenv("WEB_SOCKET_URL")
	if rbbtURL == "" {
		log.Panic("RABBITMQ_URL is not set")
	}

	conn, err := amqp.Dial(rbbtURL)
	failOnError(err, "Failed to connect to RabbitMQ")

	cws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
	failOnError(err, "Failed to connect to WebSocket")
	defer cws.Close()

	defer conn.Close()

	consumer.Consume(conn, failOnError, cws)

}
