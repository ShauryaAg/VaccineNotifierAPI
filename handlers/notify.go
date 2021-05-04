package handlers

import (
	"net/http"

	"cov-api/utils"
)

func SendNotification(w http.ResponseWriter, r *http.Request) {
	utils.SendVaccineInfo(r.Host)
}
