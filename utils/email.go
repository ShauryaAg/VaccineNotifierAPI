package utils

import (
	"fmt"

	"cov-api/models"
	"cov-api/models/db"
)

func SendConfirmationEmail(user models.User, host string) {
	var data = make(map[string]string)
	var token = models.NewUserToken(user)

	db.DBCon.Create(token) // Create and save a new token

	url := fmt.Sprintf("%s://%s/t/%s", "http", host, token.Token)
	plaincontent := fmt.Sprintf("Confirm your email by clicking on this link : %s", url)

	data["text-content"] = plaincontent
	data["html-content"] = ""
	data["to-name"] = user.Name
	data["to-email"] = user.Email
	data["from-name"] = "VaccineNotifier"
	data["from-email"] = "agora.dscbvp@gmail.com"
	data["subject"] = "Confirm Your Email!"

	SendSendgridEmail(data)
}
