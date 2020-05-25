package util

import (
	"math/rand"
	"strconv"
	"time"
)

func GetPasswordSalt() string {
	rand.Seed(time.Now().UnixNano())
	salt := strconv.Itoa(rand.Intn(8999) + 1000)
	return salt
}
