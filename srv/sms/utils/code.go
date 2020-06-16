package utils

import (
	"math/rand"
	"strconv"
	"time"
)

func GenVerificationCode() (code string) {
	code = strconv.Itoa(rand.New(rand.NewSource(time.Now().UnixNano())).Intn(899999) + 100000)
	return
}

