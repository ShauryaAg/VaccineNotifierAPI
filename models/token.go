package models

import (
	"encoding/hex"
	"fmt"
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

	fmt.Println(token.Token)

	return &token
}

func (token *Token) GenerateToken() {
	rand.Seed(time.Now().Unix())
	b := make([]byte, 10)
	if _, err := rand.Read(b); err != nil {
		return
	}
	fmt.Println(b)
	token.Token = hex.EncodeToString(b)
}
