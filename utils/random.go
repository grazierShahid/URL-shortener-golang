package utils

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateRandomKey(length int) string {
	rand.Seed(time.Now().UnixNano())
	var key []byte
	for i := 0; i < length; i++ {
		key = append(key, charset[rand.Intn(len(charset))])
	}
	return string(key)
}
