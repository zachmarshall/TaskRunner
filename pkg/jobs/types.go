package jobs

import (
	"time"
)

type JobType string

const (
	EmailJob               JobType = "Email"
	TransferStartedJob     JobType = "TransferStarted"
	TransferStateChangeJob JobType = "TransferStateChange"
	VisitStartedJob        JobType = "VisitStarted"
	VisitStateChangeJob    JobType = "VisitStateChange"
	ReminderJob            JobType = "Reminder"
)

type Job struct {
	ID          string      `json:"id"`
	Type        JobType     `json:"type"`
	Payload     interface{} `json:"payload"`
	ScheduleAt  *time.Time  `json:"schedule_at"`
	RequestedBy string      `json:"requested_by"`
	CreatedAt   time.Time   `json:"created_at"`
}

type JobHandler interface {
	Handle(job Job) error
}
