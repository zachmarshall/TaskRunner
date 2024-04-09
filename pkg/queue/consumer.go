package queue

import (
	"JobScheduler/internal/logger"
	"JobScheduler/pkg/jobs"
	"JobScheduler/pkg/jobs/dispatcher"
	"context"
	"encoding/json"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

// Consume starts consuming messages from the specified queue
func Consume(queueName string, dispatcher *dispatcher.JobDispatcher) {
	ch, err := conn.Channel()
	if err != nil {
		logger.Error("Failed to open a channel: ", err)
		return
	}
	defer ch.Close()

	if err := CreateDeadLetterExchange(ch); err != nil {
		logger.Error("Failed to create a dead-letter exchange: ", err)
		return
	}

	msgs, err := ch.Consume(
		queueName,
		"",
		true,
		false,
		false,
		false,
		nil,
	)
	if err != nil {
		logger.Error("Failed to register a consumer: ", err)
		return
	}

	forever := make(chan bool)
	jobsToProcess := make(chan jobs.Job)

	go func() {
		for d := range msgs {
			job := jobs.Job{}
			err := json.Unmarshal(d.Body, &job)
			if err != nil {
				logger.Error("Error unmarshalling message: ", err)
				continue
			}

			var delay time.Duration
			if job.ScheduleAt != nil {
				delay = time.Until(*job.ScheduleAt)
			}

			if delay > 0 {
				err = PublishDelayed("delay-queue", queueName, delay, d.Body)
				if err != nil {
					logger.Error("Error publishing delayed message: ", err)
					continue
				}
				continue
			}
			jobsToProcess <- job
		}
	}()

	go func() {
		for j := range jobsToProcess {
			newJob, err := dispatcher.Dispatch(j)
			if err != nil {
				logger.Error("Error dispatching job: ", err)
				// Publish the message to the dead-letter exchange
				ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
				defer cancel()
				err = ch.PublishWithContext(ctx,
					"dead-letter-exchange", // exchange
					"",                     // routing key
					false,                  // mandatory
					false,                  // immediate
					amqp.Publishing{
						ContentType: "text/plain",
						Body:        []byte("d.Body"),
					},
				)
				if err != nil {
					logger.Error("Failed to publish message to the dead-letter exchange: ", err)
				}
				continue
			}
			if newJob != (jobs.Job{}) {
				jobsToProcess <- newJob
			}
		}
	}()

	logger.Info("Waiting for messages. To exit press CTRL+C")
	<-forever
}
