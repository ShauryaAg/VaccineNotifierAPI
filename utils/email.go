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

func SendNotificationEmail(user models.User, AvailableSessions []interface{}) {
	var data = make(map[string]string)

	var plaincontent string
	for _, session := range AvailableSessions {
		sessionMap := session.(map[string]interface{})
		plaincontent += fmt.Sprintf("%s available at %s, %s, %s on %s for people above %d years of age\n\n", sessionMap["vaccine"], sessionMap["center"], sessionMap["district"], sessionMap["state"], sessionMap["date"], int(sessionMap["min_age_limit"].(float64)))
	}

	data["text-content"] = plaincontent
	data["html-content"] = ""
	data["to-name"] = user.Name
	data["to-email"] = user.Email
	data["from-name"] = "VaccineNotifier"
	data["from-email"] = "agora.dscbvp@gmail.com"
	data["subject"] = "Your Vaccine is Available!"

	SendSendgridEmail(data)
}
