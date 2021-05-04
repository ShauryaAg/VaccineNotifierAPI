package models

import (
	"fmt"
	"math/rand"
)

type Otp struct {
	SessionId string
	Token     string
}

func GenrateOtp() string {
	return fmt.Sprintf("%06d", rand.Intn(1000000))
}
