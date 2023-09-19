package utils

import (
	"fmt"
	"math/rand"
	"time"
)

const baseLowerCase string = "0123456789abcdefghijklmnopqrstuvwxyz"

const charset = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

func GenerateUniqueID() string {
	b := make([]byte, 4)
	for i := range b {
		b[i] = charset[seededRand.Intn(len(charset))]
	}

	return fmt.Sprint(time.Now().Unix()) + string(b)
}

func GenerateRandomId() string {
	return RandomString(32, baseLowerCase)
}

func RandomString(n int, base string) string {
	b := make([]byte, n)
	for i := range b {
		// G404: weak random number generator works for us in this scenario
		// because we need it just for task or request ids
		b[i] = base[rand.Intn(len(base))] //nolint
	}
	return string(b)
}
