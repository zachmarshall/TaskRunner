package dispatcher

import (
	"JobScheduler/pkg/email"
	"JobScheduler/pkg/jobs"
	"testing"
	"time"
)

// Send bad email job, returns true if passed, false if failed
func TestSendBadEmail(t *testing.T) {
	d := NewJobDispatcher()
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
	_, err := d.Dispatch(job)
	if err == nil {
		t.Fatal("Expected to encounter an error when sending a bad email payload type ('email')")
	}
}

func TestSendGoodEmail(t *testing.T) {
	d := NewJobDispatcher()
	j := jobs.Job{
		ID:          "2",
		Type:        "Email", // Should be capitalized
		Payload:     nil,
		ScheduleAt:  nil,
		RequestedBy: "tester",
		CreatedAt:   time.Now(),
	}

	// Send job
	_, err := d.Dispatch(j)
	if err == nil {
		t.Fatalf("Expected to encounter an error when sending a missing email payload")
	}
}
