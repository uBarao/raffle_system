package routes

import (
	"github/raffle_system/config"
	"github/raffle_system/controllers"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	r := mux.NewRouter()

	// html functions
	r.HandleFunc("/", controllers.Index)
	r.HandleFunc("/addNewBet", controllers.Feedback)
	r.HandleFunc("/addNewBet/random", controllers.FeedbackRandom)
	r.HandleFunc("/run", controllers.RunRaffleResult)

	// api functions
	r.HandleFunc("/api/addNewBet", controllers.AddNewBet).Methods("Post")
	r.HandleFunc("/api/addNewBet/random", controllers.AddNewRandomBet).Methods("Post")
	r.HandleFunc("/api/bets", controllers.AllBets).Methods("Get")
	r.HandleFunc("/api/betsByName", controllers.BetsByName).Methods("Get")
	r.HandleFunc("/api/betsByCPF", controllers.BetsByCPF).Methods("Get")
	r.HandleFunc("/api/run", controllers.RunRaffle).Methods("Get")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	log.Fatal(http.ListenAndServe(":"+config.GetPort(), handlers.CORS(handlers.AllowedOrigins([]string{"*"}))(r)))
}
