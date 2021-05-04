// https://cdn-api.co-vin.in/api/v2/appointment/sessions/public/findByPin?pincode=110054&date=06-05-2021
package handlers

import (
	"fmt"
	"net/http"
	"time"

	"cov-api/models"
	"cov-api/models/db"
	"cov-api/utils"
)

func Get(w http.ResponseWriter, r *http.Request) {
	var distinctPincode []string
	result := db.DBCon.Model(&models.User{}).Distinct("pincode").Find(&distinctPincode)
	if result.Error != nil {
		fmt.Print("err", result.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Record Found"))
		return
	}

	var users []models.User
	for _, pincode := range distinctPincode {
		db.DBCon.Find(&users, "pincode = ?", pincode)
		sessions := utils.GetVaccineDetailsByPincodeAndDate(pincode, time.Now())
		for _, user := range users {
			utils.GetAvailableSessions(user, sessions)
		}
	}

}
