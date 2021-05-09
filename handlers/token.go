package handlers

import (
	"fmt"
	"net/http"
	"time"

	"cov-api/models"
	"cov-api/models/db"

	"github.com/gorilla/mux"
)

func VerifyToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var token models.Token
	var user models.User

	db.DBCon.First(&token, "token = ?", vars["token"]).Delete(&models.Token{}) // delete token after finding
	if token.Type != "CONFIRM" || time.Now().After(token.Expiry) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid URL"))
		return
	}

	db.DBCon.First(&user, "id = ?", token.UserID)
	user.IsActive = true
	result := db.DBCon.Save(&user)
	if result.Error != nil {
		fmt.Println("err", result.Error)
		return
	}
}

func UnsubscribeToken(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var token models.Token
	var user models.User

	db.DBCon.First(&token, "token = ?", vars["token"]).Delete(&models.Token{}) // delete token after finding
	if token.Type != "UNSUBSCRIBE" || time.Now().After(token.Expiry) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid URL"))
		return
	}

	db.DBCon.First(&user, "id = ?", token.UserID)
	user.IsSubscribed = false
	result := db.DBCon.Save(&user)
	if result.Error != nil {
		fmt.Println("err", result.Error)
		return
	}
}
