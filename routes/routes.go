package routes

import (
	"github/raffle_system/config"
	"github/raffle_system/handlers"
	"github/raffle_system/utils"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()
	r.Use(utils.ContentTypeMiddleware)
	r.HandleFunc("/", handlers.Home)
	r.HandleFunc("/api/addNewBet", handlers.AddNewBet).Methods("Post")
	r.HandleFunc("/api/addNewBet/random", handlers.AddNewRandomBet).Methods("Post")
	r.HandleFunc("/api/bets", handlers.AllBets).Methods("Get")
	r.HandleFunc("/api/betsByName", handlers.BetsByName).Methods("Get")
	r.HandleFunc("/api/betsByCPF", handlers.BetsByCPF).Methods("Get")
	r.HandleFunc("/api/run", handlers.RunRaffle).Methods("Get")
	log.Fatal(http.ListenAndServe(":"+config.GetPort(), r))
}
