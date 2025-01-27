package routes

import (
	"github/raffle_system/config"
	"github/raffle_system/handlers"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	// html functions
	r.HandleFunc("/", handlers.Index)
	r.HandleFunc("/feedback", handlers.Feedback)

	// api functions
	r.HandleFunc("/api/addNewBet", handlers.AddNewBet).Methods("Post")
	r.HandleFunc("/api/addNewBet/random", handlers.AddNewRandomBet).Methods("Post")
	r.HandleFunc("/api/bets", handlers.AllBets).Methods("Get")
	r.HandleFunc("/api/betsByName", handlers.BetsByName).Methods("Get")
	r.HandleFunc("/api/betsByCPF", handlers.BetsByCPF).Methods("Get")
	r.HandleFunc("/api/run", handlers.RunRaffle).Methods("Get")
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":"+config.GetPort(), r))
}
