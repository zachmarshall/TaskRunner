package handler

import (
	"bytes"
	"fmt"
	"html/template"

	"github.com/vatusa/taskrunner/pkg/email"
	"github.com/vatusa/taskrunner/pkg/jobs"
	"github.com/vatusa/taskrunner/pkg/visit"
)

func visitEmail(visit any, tName string) (string, error) {
	t, err := template.ParseFiles(tName)
	if nil != err {
		return "", err
	}
	b := new(bytes.Buffer)
	if err = t.Execute(b, visit); err != nil {
		return "", err
	}
	return b.String(), nil
}

func HandlevisitStarted(t visit.Visit) (jobs.Job, error) {
	payload, err := visitEmail(t, "visit_started.html")
	if err != nil {
		return jobs.Job{}, err
	}
	return jobs.Job{
		Type: jobs.EmailJob,
		Payload: email.Payload{
			DestinationAddress: t.VisitorAddress,
			Subject:            "VATSIM visit Initiatied",
			Body:               payload,
		},
	}, nil
}

func HandlevisitStateChange(t visit.VisitStateChange) (jobs.Job, error) {
	var tName string
	if t.State == visit.Accepted {
		tName = "visit_accepted.html"
	} else if t.State == visit.Rejected {
		tName = "visit_rejected.html"
	} else {
		return jobs.Job{}, fmt.Errorf("unexpected visit state: %v", t.State)
	}
	payload, err := visitEmail(t, tName)
	if err != nil {
		return jobs.Job{}, err
	}
	return jobs.Job{
		Type: jobs.EmailJob,
		Payload: email.Payload{
			DestinationAddress: t.VisitorAddress,
			Subject:            fmt.Sprintf("VATSIM visit %v", t.State),
			Body:               payload,
		},
	}, nil
}
