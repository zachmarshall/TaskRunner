package handler

import (
	"JobScheduler/pkg/email"
	"JobScheduler/pkg/jobs"
	"JobScheduler/pkg/transfer"
	"bytes"
	"fmt"
	"html/template"
)

func transferEmail(transfer any, tName string) (string, error) {
	t, err := template.ParseFiles(tName)
	if nil != err {
		return "", err
	}
	b := new(bytes.Buffer)
	if err = t.Execute(b, transfer); err != nil {
		return "", err
	}
	return b.String(), nil
}

func HandleTransferStarted(t transfer.Transfer) (jobs.Job, error) {
	payload, err := transferEmail(t, "transfer_started.html")
	if err != nil {
		return jobs.Job{}, err
	}
	return jobs.Job{
		Type: jobs.EmailJob,
		Payload: email.Payload{
			DestinationAddress: t.TransfereeAddress,
			Subject:            "VATSIM Transfer Initiatied",
			Body:               payload,
		},
	}, nil
}

func HandleTransferStateChange(t transfer.TransferStateChange) (jobs.Job, error) {
	var tName string
	if t.State == transfer.Accepted {
		tName = "transfer_accepted.html"
	} else if t.State == transfer.Rejected {
		tName = "transfer_rejected.html"
	} else {
		return jobs.Job{}, fmt.Errorf("unexpected transfer state: %v", t.State)
	}
	payload, err := transferEmail(t, tName)
	if err != nil {
		return jobs.Job{}, err
	}
	return jobs.Job{
		Type: jobs.EmailJob,
		Payload: email.Payload{
			DestinationAddress: t.TransfereeAddress,
			Subject:            fmt.Sprintf("VATSIM Transfer %v", t.State),
			Body:               payload,
		},
	}, nil
}
