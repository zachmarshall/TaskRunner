package queue

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/vatusa/taskrunner/internal/logger"
)

func CreateDeadLetterExchange(ch *amqp091.Channel) error {
	// Declare a dead-letter exchange
	if err := ch.ExchangeDeclare(
		"dead-letter-exchange", // name
		"direct",               // type
		true,                   // durable
		false,                  // auto-deleted
		false,                  // internal
		false,                  // no-wait
		nil,                    // arguments
	); err != nil {
		logger.Error("Failed to declare a dead-letter exchange: ", err)
		return err
	}

	// Declare a dead-letter queue
	dlq, err := ch.QueueDeclare(
		"dead-letter-queue", // name
		true,                // durable
		false,               // delete when unused
		false,               // exclusive
		false,               // no-wait
		nil,                 // arguments
	)
	if err != nil {
		logger.Error("Failed to declare a dead-letter queue: ", err)
		return err
	}

	// Bind the dead-letter queue to the dead-letter exchange
	err = ch.QueueBind(
		dlq.Name,               // queue name
		"",                     // routing key
		"dead-letter-exchange", // exchange
		false,                  // no-wait
		nil,                    // arguments
	)
	if err != nil {
		logger.Error("Failed to bind the dead-letter queue to the dead-letter exchange: ", err)
		return err
	}

	return nil
}
