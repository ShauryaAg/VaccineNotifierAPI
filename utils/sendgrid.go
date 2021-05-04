package utils

import (
	"fmt"
	"log"
	"os"

	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

var (
	SENDGRID_API_KEY = os.Getenv("SENDGRID_API_KEY")
)

func SendSendgridEmail(data map[string]string) {
	subject := data["subject"]
	TextContent := data["text-content"]
	HtmlContent := data["html-content"]
	to := mail.NewEmail(data["to-name"], data["to-email"])
	from := mail.NewEmail(data["from-name"], data["from-email"])
	message := mail.NewSingleEmail(from, subject, to, TextContent, HtmlContent)

	client := sendgrid.NewSendClient(SENDGRID_API_KEY)

	response, err := client.Send(message)
	if err != nil {
		log.Println(err)
	} else {
		fmt.Println(response.StatusCode)
		fmt.Println(response.Body)
		fmt.Println(response.Headers)
	}
}
