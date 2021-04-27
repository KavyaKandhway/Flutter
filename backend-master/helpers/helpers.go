package helpers

import (
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
	"log"
	"strconv"
)

func SendVerificationMail(link string, loginRequest LoginParams) {
	from := mail.NewEmail("Example User", "test@furrble.com")
	subject := "Link Verificationn"
	to := mail.NewEmail("Example User", loginRequest.Email)
	plainTextContent := "Link---" + link
	htmlContent := "link " + link
	message := mail.NewSingleEmail(from, subject, to, plainTextContent, htmlContent)
	client := sendgrid.NewSendClient(SENDGRID_API_KEY)
	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		log.Println(response.StatusCode)
		log.Println(response.Body)
		log.Println(response.Headers)
	}
}