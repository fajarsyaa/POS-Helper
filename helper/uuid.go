package helper

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateRandomUUID() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(1e16)
	randomLetters := make([]rune, 3)
	for i := 0; i < 3; i++ {
		randomLetters[i] = letters[rand.Intn(len(letters))]
	}

	return fmt.Sprintf("ORD-%d-%s", randomNumber, string(randomLetters))
}

func GenerateRandomUUID2() string {
	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	randomNumber := rand.Int63n(1e16)
	randomLetters := make([]rune, 3)
	for i := 0; i < 3; i++ {
		randomLetters[i] = letters[rand.Intn(len(letters))]
	}

	return fmt.Sprintf("OII-%d-%s", randomNumber, string(randomLetters))
}
