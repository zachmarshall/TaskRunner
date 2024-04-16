package queue

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/vatusa/taskrunner/internal/logger"
	"github.com/vatusa/taskrunner/pkg/jobs"
	"github.com/vatusa/taskrunner/pkg/jobs/dispatcher"

	amqp "github.com/rabbitmq/amqp091-go"
)

func handleDispatchError(err error, j jobs.Job, ch *amqp.Channel) {
	logger.Error("Error dispatching job: ", err)
	// Try to publish the message to the dead-letter exchange
	var jBytes []byte
	jBytes, err2 := json.Marshal(j)
	if err2 != nil {
		logger.Error("Error trying to marshal job to bytes in order to publish dispatching error to rabbit: ", err2)
	} else {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		err = ch.PublishWithContext(ctx,
			"dead-letter-exchange", // exchange
			"",                     // routing key
			false,                  // mandatory
			false,                  // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        jBytes,
			},
		)
		if err != nil {
			logger.Error("Failed to publish message to the dead-letter exchange: ", err)
		}
	}
}
func SetupQueue(queueName string, ch *amqp.Channel) error {
	q, err := ch.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return err
	}
	if q.Name != queueName {
		return fmt.Errorf("created queue name '%s' did not match requested queue name '%s'", q.Name, queueName)
	}
	return nil
}

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

	if err := SetupQueue(queueName, ch); err != nil {
		logger.Error("Failed to setup rabbitmq queue:", err)
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
				handleDispatchError(err, j, ch)
			}
			if newJob != (jobs.Job{}) {
				jobsToProcess <- newJob
			}
		}
	}()

	logger.Info("Waiting for messages. To exit press CTRL+C")
	<-forever
}
