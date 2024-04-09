package dispatcher

import (
	"JobScheduler/internal/logger"
	"JobScheduler/pkg/email"
	"JobScheduler/pkg/jobs"
	"JobScheduler/pkg/transfer"
	"fmt"
)

// JobDispatcher dispatches jobs to the appropriate handlers
type JobDispatcher struct {
	// Add fields if needed, like logger, config, etc.
}

func NewJobDispatcher() *JobDispatcher {
	return &JobDispatcher{}
}

func (d *JobDispatcher) Dispatch(j jobs.Job) (jobs.Job, error) {
	logger.Info("dispatching job: ", j.ID)

	switch j.Type {
	case jobs.EmailJob:
		return d.handleEmailJob(j)
	case jobs.TransferStartedJob:
		return d.handleTransferStarted(j)
	case jobs.TransferStateChange:
		return d.handleTransferStateChange(j)

	// Add cases for other job types

	default:
		errMsg := fmt.Errorf("unsupported job type: %s", j.Type)
		logger.Error(errMsg)
		return jobs.Job{}, errMsg
	}
}

func (d *JobDispatcher) handleEmailJob(j jobs.Job) (jobs.Job, error) {
	emailJobPayload, ok := j.Payload.(email.Payload)
	if !ok {
		return jobs.Job{}, fmt.Errorf("invalid payload for email job")
	}

	return jobs.Job{}, email.SendEmail(emailJobPayload)
}

func (d *JobDispatcher) handleTransferStarted(j jobs.Job) (jobs.Job, error) {
	transferPayload, ok := j.Payload.(transfer.Transfer)
	if !ok {
		return jobs.Job{}, fmt.Errorf("invalid payload for transfer job")
	}

	return transfer.HandleTransferStarted(transferPayload)
}

func (d *JobDispatcher) handleTransferStateChange(j jobs.Job) (jobs.Job, error) {
	transferPayload, ok := j.Payload.(transfer.TransferStateChange)
	if !ok {
		return j, fmt.Errorf("invalid payload for transfer job")
	}

	return transfer.HandleTransferStateChange(transferPayload)
}
