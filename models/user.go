package models

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type preferredVaccine string

const (
	ANY         preferredVaccine = "ANY"
	COVAXIN     preferredVaccine = "COVAXIN"
	COVIDSHIELD preferredVaccine = "COVIDSHIELD"
)

type User struct {
	Id               string
	Name             string
	Password         string
	Age              int
	Pincode          string           `gorm:"index"`
	Email            string           `gorm:"index:idx_users_email,unique"`
	IsActive         bool             `sql:"DEFAULT:false"`
	IsSubscribed     bool             `sql:"DEFAULT:true"`
	PreferredVaccine preferredVaccine `sql:"DEFAULT:'ANY'"`
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
