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
	ScheduleAt  *time.Time  `json:"scheduleAt"`
	RequestedBy string      `json:"requestedBy"`
	CreatedAt   time.Time   `json:"createdAt"`
}

type JobHandler interface {
	Handle(job Job) error
}
