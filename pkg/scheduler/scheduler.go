package scheduler

import (
	"JobScheduler/internal/logger"
	"JobScheduler/pkg/jobs"
	"JobScheduler/pkg/queue"
)

type Scheduler struct {
	dispatcher  *jobs.JobDispatcher
	rabbitMQURI string
	queueName   string
}

// NewScheduler creates a new Scheduler instance
func NewScheduler(rmqURI string, queueName string) *Scheduler {
	dispatcher := jobs.NewJobDispatcher()
	return &Scheduler{
		dispatcher:  dispatcher,
		rabbitMQURI: rmqURI,
		queueName:   queueName,
	}
}

// Run starts the scheduler which in turn starts consuming messages from RabbitMQ
func (s *Scheduler) Run() {
	logger.Info("Starting scheduler")

	// Connect to RabbitMQ
	queue.Connect(s.rabbitMQURI)
	defer queue.Close()

	queue.Consume(s.queueName, s.dispatcher)
}
