package transfer

const (
	Accepted = "ACCEPTED"
	Rejected = "REJECTED"
)

// TODO: should we minmize duplicated content here and only store/expect CID from the incoming event?
type Transfer struct {
	FirstName         string `json:"fname"` // Transferee's first name
	LastName          string `json:"lname"` // Transferee's last name
	TransfereeAddress string `json:"email_address"`
	CID               string `json:"cid"`      // Transferee's CID
	FromFAC           string `json:"from_fac"` // Transferring from (FAC ID)
	ToFAC             string `json:"to_fac"`   // Transferring to (FAC ID)
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
