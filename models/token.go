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
	User   User `gorm:"constraint:OnDelete:CASCADE;"`
	UserID string
}

func NewUserToken(user User, tokenType tokenType) *Token {
	var token Token

	token.UserID = user.Id
	token.Type = tokenType
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
