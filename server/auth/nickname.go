package auth

import (
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyz"

var seededRand *rand.Rand = rand.New(rand.NewSource(time.Now().UnixNano()))

func generateRandomUserString() string {
	b := make([]int, 8)
	for i := range b {
		b[i] = seededRand.Intn(len(charset))
	}
	return string(b)
}

func GenerateUsername() string {
	return "k" + generateRandomUserString()
}
