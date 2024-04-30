package email

type Payload struct {
	DestinationAddress string   `json:"destinationAddress"` // Receiver's email address
	CCAddresses        []string `json:"ccAddresses"`        // CC email addresses
	Subject            string   `json:"subject"`            // Email subject
	Body               string   `json:"body"`               // Email body
}
