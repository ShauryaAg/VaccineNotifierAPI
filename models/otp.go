package models

import (
	"fmt"
	"math/rand"
)

type Otp struct {
	Email string
	Token string
}

func (otp *Otp) GenerateToken() {
	otp.Token = fmt.Sprintf("%010d", rand.Intn(10000000000))
}
