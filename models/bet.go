package models

import (
	"github/raffle_system/utils"
)

type Bet struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	CPF     string `json:"cpf"`
	Numbers []int8 `json:"numbers"`
}

func RandomBetNumbers(amount int) []int8 {
	newNumbers := []int8{}

	for i := 0; i < amount; i++ {
		newNumber := int8(utils.UnrepeatableRandomInt8InRange(newNumbers, 50))
		newNumbers = append(newNumbers, newNumber)
	}
	return newNumbers
}
