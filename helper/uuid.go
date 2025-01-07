package helper

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/google/uuid"
)

func GenerateRandomUUID() string {
	uuidStr := uuid.New().String()
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	numericPart := uuidStr[:16] // ambil 16 digit awal

	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	randomLetters := make([]rune, 3)
	for i := 0; i < 3; i++ {
		randomLetters[i] = letters[rand.Intn(len(letters))]
	}

	return fmt.Sprintf("ORD-%s-%s", numericPart, string(randomLetters))
}

func GenerateRandomUUID2() string {
	uuidStr := uuid.New().String()
	uuidStr = strings.ReplaceAll(uuidStr, "-", "")
	numericPart := uuidStr[:16] // ambil 16 digit awal

	letters := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	rand.Seed(time.Now().UnixNano())
	randomLetters := make([]rune, 3)
	for i := 0; i < 3; i++ {
		randomLetters[i] = letters[rand.Intn(len(letters))]
	}

	return fmt.Sprintf("OII-%s-%s", numericPart, string(randomLetters))
}
