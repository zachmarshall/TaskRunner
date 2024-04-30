package visit

const (
	Accepted = "ACCEPTED"
	Rejected = "REJECTED"
)

// TODO: should we minmize duplicated content here and only store/expect CID from the incoming event?
type Visit struct {
	FirstName      string `json:"fname"` // Visitor's first name
	LastName       string `json:"lname"` // Visitor's last name
	VisitorAddress string `json:"emailAddress"`
	CID            string `json:"cid"`     // Visitor's CID
	FromFAC        string `json:"fromFac"` // Visiting from (FAC ID)
	ToFAC          string `json:"toFac"`   // Visiting to (FAC ID)
}

// TODO: tracking these items on inbound events seems redundant. should they be stored from the initial visit event
// and fetched later (ie from a DB)? can they be looked up again on state changes (from an API)?
type VisitStateChange struct {
	Visit
	State             string   `json:"visitState"`       // The state of the visit
	Actor             string   `json:"visitActor"`       // the person who acted on the visit request
	StateChangeReason string   `json:"visitStateReason"` // the reason why the request was accepted or rejected
	Contacts          []string `json:"contacts"`         // the emails to contact in case of questions
}
