package dispatcher

import (
	"fmt"

	"github.com/vatusa/taskrunner/internal/logger"
	"github.com/vatusa/taskrunner/pkg/email"
	emailHandler "github.com/vatusa/taskrunner/pkg/email/handler"
	"github.com/vatusa/taskrunner/pkg/jobs"
	"github.com/vatusa/taskrunner/pkg/reminders"
	remindersHandler "github.com/vatusa/taskrunner/pkg/reminders/handler"
	"github.com/vatusa/taskrunner/pkg/transfer"
	transferHandler "github.com/vatusa/taskrunner/pkg/transfer/handler"
	"github.com/vatusa/taskrunner/pkg/visit"
	visitHandler "github.com/vatusa/taskrunner/pkg/visit/handler"
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
		return d.handleTransferStartedJob(j)
	case jobs.TransferStateChangeJob:
		return d.handleTransferStateChangeJob(j)
	case jobs.VisitStartedJob:
		return d.handleVisitStartedJob(j)
	case jobs.VisitStateChangeJob:
		return d.handleVisitStateChangeJob(j)
	case jobs.ReminderJob:
		return d.handleReminderJob(j)

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

	return jobs.Job{}, emailHandler.SendEmail(emailJobPayload)
}

func (d *JobDispatcher) handleTransferStartedJob(j jobs.Job) (jobs.Job, error) {
	transferPayload, ok := j.Payload.(transfer.Transfer)
	if !ok {
		return jobs.Job{}, fmt.Errorf("invalid payload for transfer job")
	}

	return transferHandler.HandleTransferStarted(transferPayload)
}

func (d *JobDispatcher) handleTransferStateChangeJob(j jobs.Job) (jobs.Job, error) {
	transferPayload, ok := j.Payload.(transfer.TransferStateChange)
	if !ok {
		return j, fmt.Errorf("invalid payload for transfer job")
	}

	return transferHandler.HandleTransferStateChange(transferPayload)
}

func (d *JobDispatcher) handleVisitStartedJob(j jobs.Job) (jobs.Job, error) {
	visitPayload, ok := j.Payload.(visit.Visit)
	if !ok {
		return j, fmt.Errorf("invalid payload for visit job")
	}

	return visitHandler.HandleVisitStarted(visitPayload)
}

func (d *JobDispatcher) handleVisitStateChangeJob(j jobs.Job) (jobs.Job, error) {
	visitPayload, ok := j.Payload.(visit.VisitStateChange)
	if !ok {
		return j, fmt.Errorf("invalid payload for visit state change job")
	}

	return visitHandler.HandleVisitStateChange(visitPayload)
}

func (d *JobDispatcher) handleReminderJob(j jobs.Job) (jobs.Job, error) {
	reminderPayload, ok := j.Payload.(reminders.Reminder)
	if !ok {
		return j, fmt.Errorf("invalid payload for reminder job")
	}

	return remindersHandler.HandleReminder(reminderPayload)
}
