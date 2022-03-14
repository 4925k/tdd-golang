package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/google/uuid"
)

// generateAlertID will generate random alertIds
func GenerateAlertID() string {

	charSet := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	part := ``
	rs := time.Now().Unix()
	rand.Seed(rs)
	for i := 0; i < 10; i++ {
		part = part + string(charSet[rand.Intn(len(charSet))])
	}
	part = fmt.Sprintf("%s_%v", part, rs)
	return part
}

func GoogleGen() string {
	return uuid.NewString()
}

func main() {
	currentID := map[string]int{}
	googleID := map[string]int{}

	for i := 0; i < 1000; i++ {
		currentID[GenerateAlertID()]++
		googleID[GoogleGen()]++
	}

	fmt.Printf("lengths\ncurrent: %d\tgoogle: %d\n", len(currentID), len(googleID))
	fmt.Println(currentID)
	fmt.Println(googleID)
}
