package utils

import (
	"fmt"

	"cov-api/models"
	"cov-api/models/db"
)

func SendConfirmationEmail(user models.User, host string) error {
	var data = make(map[string]string)
	var token = models.NewUserToken(user, "CONFIRM")

	db.DBCon.Create(token) // Create and save a new token

	url := fmt.Sprintf("%s://%s/t/%s", "http", host, token.Token)
	plaincontent := fmt.Sprintf("Confirm your email by clicking on this link : %s", url)

	data["text-content"] = plaincontent

	data["html-content"] = ParseTemplate("templates/emails/confirm_email.html", struct {
		Name string
		Url  string
	}{user.Name, url})

	data["to-name"] = user.Name
	data["to-email"] = user.Email
	data["from-name"] = "VaccineNotifier"
	data["from-email"] = "agora.dscbvp@gmail.com"
	data["subject"] = "Confirm Your Email!"

	return SendSendgridEmail(data)
}

func SendNotificationEmail(user models.User, host string, AvailableSessions []interface{}) error {
	var data = make(map[string]string)

	var token = models.NewUserToken(user, "UNSUBSCRIBE")

	db.DBCon.Create(token) // Create and save a new token

	var plaincontent string
	for _, session := range AvailableSessions {
		sessionMap := session.(map[string]interface{})
		plaincontent += fmt.Sprintf("%s available at %s, %s, %s on %s for people above %d years of age\n\n", sessionMap["vaccine"], sessionMap["center"], sessionMap["district"], sessionMap["state"], sessionMap["date"], int(sessionMap["min_age_limit"].(float64)))
	}

	url := fmt.Sprintf("%s://%s/u/%s", "http", host, token.Token)
	plaincontent += fmt.Sprintf("\n Unsubscribe from further emails using this link: %s", url)

	data["text-content"] = plaincontent

	data["html-content"] = ParseTemplate("templates/emails/notification.html", struct {
		Sessions []interface{}
		Url      string
	}{AvailableSessions, url})

	data["to-name"] = user.Name
	data["to-email"] = user.Email
	data["from-name"] = "VaccineNotifier"
	data["from-email"] = "agora.dscbvp@gmail.com"
	data["subject"] = "Your Vaccine is Available!"

	return SendSendgridEmail(data)
}

func SendPasswordResetEmail(user models.User, host string) error {
	var data = make(map[string]string)

	var token = models.NewUserToken(user, "FORGOT")

	db.DBCon.Create(token) // Create and save a new token

	url := fmt.Sprintf("%s://%s/f/%s", "http", host, token.Token)
	plaincontent := fmt.Sprintf("Reset your password using this link : %s\n\n Valid only for 30 Minutes", url)

	data["text-content"] = plaincontent

	data["html-content"] = ParseTemplate("templates/emails/reset_password.html", struct {
		Name string
		Url  string
	}{user.Name, url})

	data["to-name"] = user.Name
	data["to-email"] = user.Email
	data["from-name"] = "VaccineNotifier"
	data["from-email"] = "agora.dscbvp@gmail.com"
	data["subject"] = "Reset Your Password!"

	return SendSendgridEmail(data)
}
