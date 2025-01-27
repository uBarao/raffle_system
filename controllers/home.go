package controllers

import (
	"encoding/json"
	"github/raffle_system/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "Index", nil)
}

func Feedback(w http.ResponseWriter, r *http.Request) {
	var newBet models.Bet
	json.NewDecoder(r.Body).Decode(&newBet)
	newBet, err := models.ActiveRaffle.AddBet(newBet.Name, newBet.CPF, newBet.Numbers)
	if err != nil {
		http.Error(w, "Error in bet", http.StatusBadRequest)
	} else {
		temp.ExecuteTemplate(w, "Feedback", newBet)
	}
}

func FeedbackRandom(w http.ResponseWriter, r *http.Request) {
	var newBet models.Bet
	json.NewDecoder(r.Body).Decode(&newBet)
	newBet.Numbers = models.RandomBetNumbers(5)
	newBet, err := models.ActiveRaffle.AddBet(newBet.Name, newBet.CPF, newBet.Numbers)
	if err != nil {
		http.Error(w, "Error in bet", http.StatusBadRequest)
	} else {
		temp.ExecuteTemplate(w, "Feedback", newBet)
	}
}

func RunRaffleResult(w http.ResponseWriter, r *http.Request) {
	if len(models.ActiveRaffle.Bets) <= 0 {
		http.Error(w, "No bets counted.", http.StatusBadRequest)
	}
	sortedNumbers, winningBets := models.ActiveRaffle.Run()
	result := models.CalculateResult(sortedNumbers, winningBets, models.ActiveRaffle.Bets)
	temp.ExecuteTemplate(w, "Result", result)
}
