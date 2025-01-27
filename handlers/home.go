package handlers

import (
	"github/raffle_system/models"
	"html/template"
	"net/http"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	//w.Header().Set("Content-type", "text/html; charset=utf-8")
	temp.ExecuteTemplate(w, "Index", nil)
}

func Feedback(w http.ResponseWriter, r *http.Request) {
	bet := models.Bet{Id: 1, Name: "Sim", CPF: "talvez", Numbers: []int8{1, 4, 6, 43, 43}}
	temp.ExecuteTemplate(w, "Feedback", bet)
}
