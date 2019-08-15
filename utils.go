package main

import (
	"math/rand"
	"time"
)

var chars = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func generateID(length int) string {
	rand.Seed(time.Now().UnixNano())
	id := make([]rune, length)
	for i := range id {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return string(id)
}
