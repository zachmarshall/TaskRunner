package handler

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/vatusa/taskrunner/pkg/email"
	"github.com/vatusa/taskrunner/pkg/jobs"
	"github.com/vatusa/taskrunner/pkg/reminders"
)

func reminderEmail(reminder any, tName string) (string, error) {
	t, err := template.ParseFiles(tName)
	if nil != err {
		return "", err
	}
	b := new(bytes.Buffer)
	if err = t.Execute(b, reminder); err != nil {
		return "", err
	}
	return b.String(), nil
}

func HandleReminder(r reminders.Reminder) (jobs.Job, error) {
	payload, err := reminderEmail(r, "reminder.html")
	if err != nil {
		return jobs.Job{}, err
	}
	return jobs.Job{
		Type: jobs.EmailJob,
		Payload: email.Payload{
			DestinationAddress: "", //TODO: who does this get sent to? where do we get that email?
			Subject:            fmt.Sprintf("VATSIM: Pending %v requires action", r.ActionType),
			Body:               payload,
		},
	}, nil
}
