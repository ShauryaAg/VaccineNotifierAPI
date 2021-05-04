package handlers

import (
	"net/http"

	"cov-api/utils"
)

func Get(w http.ResponseWriter, r *http.Request) {
	utils.SendVaccineInfo()
}
