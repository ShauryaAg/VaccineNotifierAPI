package views

import (
	"cov-api/models"
	"cov-api/models/db"
	"fmt"
	"html/template"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
)

func ResetPasswordView(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var token models.Token

	db.DBCon.First(&token, "token = ?", vars["token"])
	if token.Type != "FORGOT" || time.Now().After(token.Expiry) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid URL"))
		return
	}

	wd, err := os.Getwd()
	if err != nil {
		fmt.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	t, err := template.ParseFiles(wd + "/templates/password_reset.html")
	if err != nil {
		fmt.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}

	err = t.Execute(w, nil)
	if err != nil {
		fmt.Println("error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("Internal Server Error"))
		return
	}
}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var token models.Token
	var user models.User

	r.ParseForm() // parsing application/x-www-form-urlencoded

	db.DBCon.First(&token, "token = ?", vars["token"]).Delete(&models.Token{}) // delete token after finding
	if token.Type != "FORGOT" || time.Now().After(token.Expiry) {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid URL"))
		return
	}

	db.DBCon.First(&user, "id = ?", token.UserID)
	user.SetPassword(r.Form["password"][0])
	result := db.DBCon.Save(&user)
	if result.Error != nil {
		fmt.Println("err", result.Error)
		return
	}
}
