package reminders

type Reminder struct {
	FirstName              string `json:"fname"`         // Visitor's first name
	LastName               string `json:"lname"`         // Visitor's last name
	ControllerEmailAddress string `json:"email_address"` // email address of controller who this reminder is about
	CID                    string `json:"cid"`           // Visitor's CID
	FromFAC                string `json:"from_fac"`      // Visiting from (FAC ID)
	ToFAC                  string `json:"to_fac"`        // Visiting to (FAC ID)
	ActionType             string `json:"action_type"`   // the action (visit or transfer)
	ActionAgeInDays        int    `json:"action_age"`    // the age of the pending action
}
