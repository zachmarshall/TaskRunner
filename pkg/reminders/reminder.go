package reminders

type Reminder struct {
	FirstName              string `json:"firstName"`    // Visitor's first name
	LastName               string `json:"lastName"`     // Visitor's last name
	ControllerEmailAddress string `json:"emailAddress"` // email address of controller who this reminder is about
	CID                    string `json:"cid"`          // Visitor's CID
	FromFAC                string `json:"fromFac"`      // Visiting from (FAC ID)
	ToFAC                  string `json:"toFac"`        // Visiting to (FAC ID)
	ActionType             string `json:"actionType"`   // the action (visit or transfer)
	ActionAgeInDays        int    `json:"actionAge"`    // the age of the pending action
}
