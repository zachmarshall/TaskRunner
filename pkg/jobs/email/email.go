package email

type Payload struct {
	DestinationAddress string   `json:"destination_address"` // Receiver's email address
	CCAddresses        []string `json:"cc_addresses"`        // CC email addresses
	Subject            string   `json:"subject"`             // Email subject
	Body               string   `json:"body"`                // Email body
}
