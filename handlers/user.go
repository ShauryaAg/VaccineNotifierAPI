package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"cov-api/models"
	"cov-api/models/db"
	"cov-api/utils"

	"github.com/google/uuid"
)

func Register(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("Need content-type: 'application/json', but got %s", ct)))
		return
	}

	var user models.User
	err = json.Unmarshal(bodyBytes, &user)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	user.Id = uuid.New().String()
	user.IsSubscribed = true
	user.HashPassword()
	db.DBCon.Create(&user)

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func Login(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if ct != "application/json" {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("Need content-type: 'application/json', but got %s", ct)))
		return
	}

	var data map[string]string
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	var user models.User
	db.DBCon.First(&user, "email = ?", data["email"])

	valid := user.VerifyPassword(data["password"])
	var token string
	if valid {
		token, err = utils.CreateToken(user)
		if err != nil {
			fmt.Println(err)
		}
	}

	jsonBytes, err := json.Marshal(struct {
		Id    string
		Email string
		Token string
	}{user.Id, user.Email, token})
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

// Get user details using JWT
func GetUser(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("decoded")

	var user models.User
	db.DBCon.First(&user, "email = ?", email)

	jsonBytes, err := json.Marshal(user)

	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}
