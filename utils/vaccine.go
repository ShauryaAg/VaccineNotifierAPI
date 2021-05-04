package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"cov-api/models"
)

func GetVaccineDetailsByPincodeAndDate(pincode string, date time.Time) map[string]interface{} {
	url := fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByPin?pincode=%s&date=%s", pincode, date.Format("02-01-2006"))
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("err", err)
	}
	bytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var sessions map[string]interface{}
	err = json.Unmarshal(bytes, &sessions)
	if err != nil {
		fmt.Println("err", err)
	}

	return sessions
}

func GetAvailableSessions(user models.User, sessions map[string]interface{}) []interface{} {

	var AvailableSessions []interface{}
	for _, session := range sessions["sessions"].([]interface{}) {
		sessionMap := session.(map[string]interface{})
		if int(sessionMap["min_age_limit"].(float64)) < user.Age && (sessionMap["vaccine"] == user.PreferredVaccine || sessionMap["vaccine"] == "ANY") {
			AvailableSessions = append(AvailableSessions, session)
		}
	}

	return AvailableSessions
}
