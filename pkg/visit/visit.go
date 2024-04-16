package visit

const (
	Accepted = "ACCEPTED"
	Rejected = "REJECTED"
)

// TODO: should we minmize duplicated content here and only store/expect CID from the incoming event?
type Visit struct {
	FirstName      string `json:"fname"` // Visitor's first name
	LastName       string `json:"lname"` // Visitor's last name
	VisitorAddress string `json:"email_address"`
	CID            string `json:"cid"`      // Visitor's CID
	FromFAC        string `json:"from_fac"` // Visiting from (FAC ID)
	ToFAC          string `json:"to_fac"`   // Visiting to (FAC ID)
}

// TODO: tracking these items on inbound events seems redundant. should they be stored from the initial visit event
// and fetched later (ie from a DB)? can they be looked up again on state changes (from an API)?
type VisitStateChange struct {
	Visit
	State             string   `json:"visit_state"`        // The state of the visit
	Actor             string   `json:"visit_actor"`        // the person who acted on the visit request
	StateChangeReason string   `json:"visit_state_reason"` // the reason why the request was accepted or rejected
	Contacts          []string `json:"contacts"`           // the emails to contact in case of questions
}
