package main

import (
	"math/rand"
	"time"
)

var chars = []rune("23456789abcdefghkmnopqrstuvwxyzABCDEFGHJKLMNPQRSTUVWXYZ")

func generateID(length int) string {
	rand.Seed(time.Now().UnixNano())
	id := make([]rune, length)
	for i := range id {
		id[i] = chars[rand.Intn(len(chars))]
	}
	return string(id)
}
