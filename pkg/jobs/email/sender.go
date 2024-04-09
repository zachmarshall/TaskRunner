package email

import (
	"fmt"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

func SendEmail(payload Payload) error {
	from := mail.NewEmail("VATUSA Mailman", "mailman@vatusa.net")
	subject := payload.Subject
	to := mail.NewEmail("", payload.DestinationAddress)
	plainTextContent := payload.Body
	htmlContent := "<p>" + payload.Body + "</p>"
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)

	if len(payload.CCAddresses) > 0 {
		ccEmails := make([]*mail.Email, len(payload.CCAddresses))
		for i, cc := range payload.CCAddresses {
			ccEmails[i] = mail.NewEmail("", cc)
		}

		message.Personalizations[0].AddCCs(ccEmails...)
	}

	client := sendgrid.NewSendClient(os.Getenv("SENDGRID_API_KEY"))
	response, err := client.Send(message)
	if err != nil {
		return nil, err
	}

	if response.StatusCode == 200 {
		return nil, nil
	}

	return nil, fmt.Errorf("failed to send email, status code: %d", response.StatusCode)
}
