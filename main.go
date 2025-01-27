package main

import (
	"github/raffle_system/models"
	"github/raffle_system/routes"
)

func main() {
	models.Start()
	routes.HandleRequest()
}
