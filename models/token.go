package models

import (
	"encoding/hex"
	"math/rand"
	"time"
)

type Token struct {
	Token  string
	User   User `gorm:"constraint:OnDelete:CASCADE;"`
	UserID string
}

func NewUserToken(user User) *Token {
	var token Token

	token.UserID = user.Id
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
