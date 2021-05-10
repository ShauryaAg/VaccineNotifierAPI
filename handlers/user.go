package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

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
	if !strings.Contains(ct, "application/json") {
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
	result := db.DBCon.Create(&user)
	if result.Error != nil {
		if strings.Contains(
			result.Error.Error(),
			"ERROR: duplicate key value violates unique constraint",
		) {
			user.Id = "" // a bit of a hack to make userId nil for searching
			result := db.DBCon.First(&user, "email = ?", user.Email)
			if result.Error != nil {
				fmt.Println("err", result.Error)
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte("Internal Server Error"))
				return
			}
			if !user.IsActive {
				goto sendMail // Send mail again even if user exists
			} else {
				w.WriteHeader(http.StatusBadRequest)
				w.Write([]byte("Email already in use"))
				return
			}
		}
		fmt.Println("err", result.Error)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(result.Error.Error()))
		return
	}

sendMail:
	err = utils.SendConfirmationEmail(user, r.Host) // Sending Confirmation Email
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Please confirm your Email"))
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
	if !strings.Contains(ct, "application/json") {
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
	result := db.DBCon.First(&user, "email = ?", data["email"])
	if result.Error != nil {
		fmt.Print("err", result.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(result.Error.Error()))
		return
	} else if !user.IsActive {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please confirm your Email!"))
		return
	}

	valid := user.VerifyPassword(data["password"])
	var token string
	if valid {
		token, err = utils.CreateToken(user)
		if err != nil {
			fmt.Println(err)
			return
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Email/Password is incorrect"))
		return
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
		return
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
	if !user.IsActive {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please confirm your Email!"))
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("decoded")

	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if !strings.Contains(ct, "application/json") {
		w.WriteHeader(http.StatusUnsupportedMediaType)
		w.Write([]byte(fmt.Sprintf("Need content-type: 'application/json', but got %s", ct)))
		return
	}

	var data map[string]interface{}
	err = json.Unmarshal(bodyBytes, &data)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	var user models.User
	// woo bad hack
	if data["Password"] != nil || data["password"] != nil {
		var password string
		if data["Password"] != nil {
			password = data["Password"].(string)
		} else {
			password = data["password"].(string)
		}

		user.SetPassword(password)

		if data["Password"] != nil {
			data["Password"] = user.Password
		} else {
			data["password"] = user.Password
		}
	}
	result := db.DBCon.Model(&models.User{}).Where("email = ?", email).Updates(data).First(&user) // Updates and stores it in &user
	if result.Error != nil {
		fmt.Print("err", result.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("No Record Found"))
		return
	}

	var jsonBytes []byte
	jsonBytes, err = json.Marshal(user)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func UnsubscribeUser(w http.ResponseWriter, r *http.Request) {
	email := r.Header.Get("decoded")

	var user models.User
	db.DBCon.First(&user, "email = ?", email)
	if !user.IsActive {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("Please confirm your Email!"))
		return
	}

	user.IsSubscribed = false
	result := db.DBCon.Save(&user)
	if result.Error != nil {
		fmt.Println("err", result.Error)
		return
	}

	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonBytes)
}

func ResetUserPassword(w http.ResponseWriter, r *http.Request) {
	bodyBytes, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		fmt.Print("err", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	ct := r.Header.Get("content-type")
	if !strings.Contains(ct, "application/json") {
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
	result := db.DBCon.First(&user, "email = ?", data["email"])
	if result.Error != nil || !user.IsActive {
		fmt.Print("err", result.Error)
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Invalid Email"))
		return
	}

	utils.SendPasswordResetEmail(user, r.Host) // send reset email

	w.Header().Add("content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Password Reset Email has been sent"))
}
