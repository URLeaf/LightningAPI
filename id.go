package main

import (
	"math/rand"
	"time"
)

func makeID(length int) string {
	rand.Seed(time.Now().Unix())

	var result string
	var characters = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"
	for i := 0; i < length; i++ {
		result += string(characters[rand.Intn(len(characters))])
	}
	return result
}
