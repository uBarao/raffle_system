package handlers

import (
	"encoding/json"
	"github/raffle_system/models"
	"net/http"
)

func Start() {
	models.ActiveRaffle = *models.Start()
}

func AddNewBet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var newBet models.Bet
	json.NewDecoder(r.Body).Decode(&newBet)
	newBet, err := models.ActiveRaffle.AddBet(newBet.Name, newBet.CPF, newBet.Numbers)
	if err != nil {
		http.Error(w, "Error in bet", http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(newBet)
	}
}

func AddNewRandomBet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var newBet models.Bet
	json.NewDecoder(r.Body).Decode(&newBet)
	newBet.Numbers = models.RandomBetNumbers(5)
	newBet, err := models.ActiveRaffle.AddBet(newBet.Name, newBet.CPF, newBet.Numbers)
	if err != nil {
		http.Error(w, "Error in bet", http.StatusBadRequest)
	} else {
		json.NewEncoder(w).Encode(newBet)
	}
}

func AllBets(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(models.ActiveRaffle.Bets)
}

func BetsByName(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var nameToFind string
	json.NewDecoder(r.Body).Decode(&nameToFind)
	bets := models.ActiveRaffle.FindBetByName(nameToFind)
	json.NewEncoder(w).Encode(bets)
}

func BetsByCPF(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var cpfToFind string
	json.NewDecoder(r.Body).Decode(&cpfToFind)
	bets := models.ActiveRaffle.FindBetByCPF(cpfToFind)
	json.NewEncoder(w).Encode(bets)
}

func RunRaffle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	sortedNumbers, winningBets := models.ActiveRaffle.Run()
	result := models.CalculateResult(sortedNumbers, winningBets)
	json.NewEncoder(w).Encode(result)
}
