package test

import (
	"JobScheduler/pkg/jobs"
	"JobScheduler/pkg/jobs/email"
	"time"
)

// Send bad email job, returns true if passed, false if failed
func SendBadEmail(d *jobs.JobDispatcher) bool {
	// Create job payload
	job := jobs.Job{
		ID:   "1",
		Type: "email", // Should be capitalized
		Payload: email.Payload{
			DestinationAddress: "vatusa6@vatusa.net",
			CCAddresses:        []string{"vatusa4@vatusa.net"},
			Subject:            "Test email",
			Body:               "This is a test email",
		},
		ScheduleAt:  nil,
		RequestedBy: "tester",
		CreatedAt:   time.Now(),
	}

	// Send job
	err := d.Dispatch(job)
	if err == nil {
		return false
	}

	job = jobs.Job{
		ID:          "2",
		Type:        "Email", // Should be capitalized
		Payload:     nil,
		ScheduleAt:  nil,
		RequestedBy: "tester",
		CreatedAt:   time.Now(),
	}

	// Send job
	err = d.Dispatch(job)
	if err == nil {
		return false
	}

	return true
}
