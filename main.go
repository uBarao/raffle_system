package main

import (
	"github/raffle_system/models"
	"github/raffle_system/routes"
)

func main() {
	raffle := models.Start()

	raffle.AddBet("Robson", "12345678910", models.RandomBetNumbers(5))
	raffle.AddBet("Robson", "12345678910", models.RandomBetNumbers(5))
	raffle.AddBet("Robson", "12345678910", models.RandomBetNumbers(5))

	routes.HandleRequest()
}
