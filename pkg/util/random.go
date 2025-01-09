package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.New(rand.NewSource(time.Now().UnixNano()))
}

func GenerateRandomString(n int) string {
	var charsets = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	letter := make([]rune, n)
	for i := range letter {
		letter[i] = charsets[rand.Intn(len(charsets))]
	}

	return string(letter)
}

func GenerateRandomNumber(n int) string {
	var charsets = []rune("0123456789")
	letter := make([]rune, n)
	for i := range letter {
		letter[i] = charsets[rand.Intn(len(charsets))]
	}

	return string(letter)
}
