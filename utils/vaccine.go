package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"cov-api/models"
	"cov-api/models/db"
)

func GetVaccineDetailsByPincodeAndDate(pincode string, date time.Time) map[string]interface{} {
	url := fmt.Sprintf("https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/calendarByPin?pincode=%s&date=%s", pincode, date.Format("02-01-2006"))
	response, err := http.Get(url)
	if err != nil {
		fmt.Println("err", err)
	}
	bytes, _ := ioutil.ReadAll(response.Body)
	defer response.Body.Close()

	var centers map[string]interface{}
	err = json.Unmarshal(bytes, &centers)
	if err != nil {
		fmt.Println("err", err)
	}

	return centers
}

func GetAvailableSessions(user models.User, centers map[string]interface{}) []interface{} {

	var AvailableSessions []interface{}

	for _, center := range centers["centers"].([]interface{}) {
		centerMap := center.(map[string]interface{})
		for _, session := range centerMap["sessions"].([]interface{}) {
			sessionMap := session.(map[string]interface{})

			sessionMap["center"] = centerMap["address"]
			sessionMap["district"] = centerMap["district_name"]
			sessionMap["state"] = centerMap["state_name"]

			if int(sessionMap["min_age_limit"].(float64)) < user.Age && (strings.Compare(sessionMap["vaccine"].(string), string(user.PreferredVaccine)) == 0 || strings.Compare(sessionMap["vaccine"].(string), "ANY") == 0) {
				AvailableSessions = append(AvailableSessions, sessionMap)
			}
		}
	}

	return AvailableSessions
}

func SendVaccineInfo(host string) {
	var distinctPincode []string
	result := db.DBCon.Model(&models.User{}).Distinct("pincode").Find(&distinctPincode)
	if result.Error != nil {
		fmt.Print("err", result.Error)
		return
	}

	var users []models.User
	for _, pincode := range distinctPincode {
		db.DBCon.Find(&users, "pincode = ?", pincode)
		centers := GetVaccineDetailsByPincodeAndDate(pincode, time.Now())
		for _, user := range users {
			if user.IsActive && user.IsSubscribed {
				AvailableSessions := GetAvailableSessions(user, centers)
				if AvailableSessions != nil {
					SendNotificationEmail(user, host, AvailableSessions)
				}
			}
		}
	}
}
