package transfer

const (
	Accepted = "ACCEPTED"
	Rejected = "REJECTED"
)

// TODO: should we minmize duplicated content here and only store/expect CID from the incoming event?
type Transfer struct {
	FirstName         string `json:"firstName"` // Transferee's first name
	LastName          string `json:"lastName"`  // Transferee's last name
	TransfereeAddress string `json:"emailAddress"`
	CID               string `json:"cid"`     // Transferee's CID
	FromFAC           string `json:"fromFac"` // Transferring from (FAC ID)
	ToFAC             string `json:"toFac"`   // Transferring to (FAC ID)
	Reason            string `json:"transferReason"`
}

// TODO: tracking these items on inbound events seems redundant. should they be stored from the initial transfer event
// and fetched later (ie from a DB)? can they be looked up again on state changes (from an API)?
type TransferStateChange struct {
	Transfer
	State             string   `json:"transferState"`       // The state of the transfer
	Actor             string   `json:"transferActor"`       // the person who acted on the transfer request
	StateChangeReason string   `json:"transferStateReason"` // the reason why the request was accepted or rejected
	Contacts          []string `json:"contacts"`            // the emails to contact in case of questions
}
