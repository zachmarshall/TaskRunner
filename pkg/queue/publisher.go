package queue

import (
	"JobScheduler/internal/logger"
	"context"
	amqp "github.com/rabbitmq/amqp091-go"
	"time"
)

// Publish publishes a job message to the specified queue
func Publish(queueName string, jobMessage []byte) error {
	ch, err := conn.Channel()
	if err != nil {
		logger.Error("Failed to open a channel: ", err)
		return err
	}
	defer ch.Close()

	// Ensure the queue exists
	_, err = ch.QueueDeclare(
		queueName,
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Error("Failed to declare a queue: ", err)
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Publish the message
	err = ch.PublishWithContext(ctx,
		"",        // Exchange
		queueName, // Routing key (queue name)
		false,     // Mandatory
		false,     // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Body:        jobMessage,
		},
	)
	if err != nil {
		logger.Error("Failed to publish a message: ", err)
		return err
	}

	return nil
}

// PublishDelayed publishes a job message to a delayed exchange with a specified delay
func PublishDelayed(exchangeName, routingKey string, delay time.Duration, jobMessage []byte) error {
	ch, err := conn.Channel()
	if err != nil {
		logger.Error("Failed to open a channel: ", err)
		return err
	}
	defer ch.Close()

	// Declare a delayed exchange
	err = ch.ExchangeDeclare(
		exchangeName,                           // Exchange name
		"x-delayed-message",                    // Exchange type (delayed)
		true,                                   // Durable
		false,                                  // Auto-deleted
		false,                                  // Internal
		false,                                  // No-wait
		amqp.Table{"x-delayed-type": "direct"}, // Arguments
	)
	if err != nil {
		logger.Error("Failed to declare a delayed exchange: ", err)
		return err
	}

	// Publish the message with a delay
	headers := amqp.Table{}
	if delay > 0 {
		headers["x-delay"] = int32(delay.Milliseconds())
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		exchangeName, // Exchange
		routingKey,   // Routing key
		false,        // Mandatory
		false,        // Immediate
		amqp.Publishing{
			ContentType: "application/json",
			Headers:     headers,
			Body:        jobMessage,
		},
	)
	if err != nil {
		logger.Error("Failed to publish a delayed message: ", err)
		return err
	}

	return nil
}
