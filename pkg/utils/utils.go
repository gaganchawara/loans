package utils

import "math/rand"

const baseLowerCase string = "0123456789abcdefghijklmnopqrstuvwxyz"

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
