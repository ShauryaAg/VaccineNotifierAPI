package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type preferredVaccine int

const (
	ANY         preferredVaccine = iota // 0
	COVAXIN                             // 1
	COVIDSHIELD                         // 2
)

type User struct {
	Id               string
	Email            string
	Name             string
	Age              int
	Password         string
	Pincode          string
	IsSubscribed     bool
	PreferredVaccine preferredVaccine
}

func (a *User) HashPassword() {
	hash, err := bcrypt.GenerateFromPassword([]byte(a.Password), bcrypt.DefaultCost)
	if err != nil {
		fmt.Println(err)
		return
	}

	a.Password = string(hash)
}

func (a *User) VerifyPassword(attempt string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(a.Password), []byte(attempt))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
