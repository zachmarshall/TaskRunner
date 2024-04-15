package queue

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/vatusa/taskrunner/internal/logger"
)

var conn *amqp.Connection

// Connect initializes a connection to RabbitMQ
func Connect(uri string) {
	var err error

	for {
		conn, err = amqp.Dial(uri)
		if err != nil {
			logger.Error("Failed to connect to RabbitMQ: ", err)
			log.Println("Retrying in 5 seconds...")
			time.Sleep(5 * time.Second)
			continue
		}

		logger.Info("Connected to RabbitMQ")

		// Listen for a disconnection event
		go func() {
			closeErr := <-conn.NotifyClose(make(chan *amqp.Error))
			if closeErr != nil {
				logger.Error("RabbitMQ connection closed, reconnecting: ", closeErr)
				Connect(uri)
			}
		}()

		break
	}
}

// Close closes the RabbitMQ connection
func Close() {
	if conn != nil {
		err := conn.Close()
		if err != nil {
			logger.Error("Failed to close RabbitMQ connection: ", err)
		}
	}
}
