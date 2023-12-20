package jobs

import (
	"JobScheduler/internal/logger"
	"JobScheduler/pkg/jobs/email"
	"fmt"
)

// JobDispatcher dispatches jobs to the appropriate handlers
type JobDispatcher struct {
	// Add fields if needed, like logger, config, etc.
}

func NewJobDispatcher() *JobDispatcher {
	return &JobDispatcher{}
}

func (d *JobDispatcher) Dispatch(job Job) error {
	logger.Info("dispatching job: ", job.ID)

	switch job.Type {
	case EmailJob:
		return d.handleEmailJob(job)

	// Add cases for other job types

	default:
		errMsg := fmt.Errorf("unsupported job type: %s", job.Type)
		logger.Error(errMsg)
		return errMsg
	}
}

func (d *JobDispatcher) handleEmailJob(job Job) error {
	emailJobPayload, ok := job.Payload.(email.Payload)
	if !ok {
		return fmt.Errorf("invalid payload for email job")
	}

	return email.SendEmail(emailJobPayload)
}
