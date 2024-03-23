package transfer

import (
	"JobScheduler/pkg/jobs/email"
	"bytes"
	"fmt"
	"html/template"
)

const (
	Accepted = "ACCEPTED"
	Rejected = "REJECTED"
)

// TODO: should we minmize duplicated content here and only store/expect CID from the incoming event?
type Transfer struct {
	FirstName         string `json:"fname"` // Transferee's first name
	LastName          string `json:"lname"` // Transferee's last name
	TransfereeAddress string `json:"email_address"`
	CID               string `json:"cid"`           // Transferee's CID
	TransferFrom      string `json:"transfer_from"` // Transferring from (FAC ID)
	TransferTo        string `json:"transfer_to"`   // Transferring to (FAC ID)
	Reason            string `json:"transfer_reason"`
}

// TODO: tracking these items on inbound events seems redundant. should they be stored from the initial transfer event
// and fetched later (ie from a DB)? can they be looked up again on state changes (from an API)?
type TransferStateChange struct {
	Transfer
	State             string   `json:"transfer_state"`        // The state of the transfer
	Actor             string   `json:"transfer_actor"`        // the person who acted on the transfer request
	StateChangeReason string   `json:"transfer_state_reason"` // the reason why the request was accepted or rejected
	Contacts          []string `json:"contacts"`              // the emails to contact in case of questions
}

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

func HandleTransferStarted(transfer Transfer) error {
	payload, err := transferEmail(transfer, "transfer_started.html")
	if err != nil {
		return err
	}
	email.SendEmail(email.Payload{
		DestinationAddress: transfer.TransfereeAddress,
		Subject:            "VATSIM Transfer Initiatied",
		Body:               payload,
	})
	return nil
}

func HandleTransferStateChange(transfer TransferStateChange) error {
	var tName string
	if transfer.State == Accepted {
		tName = "transfer_accepted.html"
	} else if transfer.State == Rejected {
		tName = "transfer_rejected.html"
	} else {
		return fmt.Errorf("unexpected transfer state: %v", transfer.State)
	}
	payload, err := transferEmail(transfer, tName)
	if err != nil {
		return err
	}
	email.SendEmail(email.Payload{
		DestinationAddress: transfer.TransfereeAddress,
		Subject:            "VATSIM Transfer Initiatied",
		Body:               payload,
	})
	return nil
}
