package models

import (
	"encoding/hex"
	"math/rand"
	"time"
)

type tokenType string

const (
	CONFIRM     tokenType = "CONFIRM"
	UNSUBSCRIBE tokenType = "UNSUBSCRIBE"
	FORGOT      tokenType = "FORGOT"
)

type Token struct {
	Token  string
	Type   tokenType
	Expiry time.Time
	UserID string
	User   User `gorm:"constraint:OnDelete:CASCADE;"`
}

func NewUserToken(user User, tokenType tokenType) *Token {
	var token Token

	token.UserID = user.Id
	token.Type = tokenType
	if tokenType == "FORGOT" {
		token.Expiry = time.Now().Add(30 * time.Minute) // valid for 30 minutes only
	} else {
		token.Expiry = time.Now().AddDate(10, 0, 0) // Valid for 10 years from now
	}
	token.GenerateToken()

	return &token
}

func (token *Token) GenerateToken() {
	rand.Seed(time.Now().Unix())
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		return
	}
	token.Token = hex.EncodeToString(b)
}
