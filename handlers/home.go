package handlers

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
