package utils

import (
	"math/rand"
	"time"
)

func UnrepeatableRandomInt8InRange(presentNumbers []int8, interval int8) int8 {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	newNumber := r.Intn(int(interval)-1) + 1

	for alreadyPresent(int8(newNumber), presentNumbers) {
		newNumber = r.Intn(int(interval)-1) + 1
	}
	return int8(newNumber)
}

func alreadyPresent(toFind int8, numbers []int8) bool {
	for _, n := range numbers {
		if toFind == int8(n) {
			return true
		}
	}
	return false
}
