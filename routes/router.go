package routes

import (
	"github.com/gorilla/mux"
	"naratmalsami-survey-server/db"
	"naratmalsami-survey-server/service"
)

func SetupRouter(wordDB *db.DataDB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/voted", service.InsertVotedRating(wordDB)).Methods("POST")

	r.HandleFunc("/sheet", service.GetSheetService(wordDB)).Methods("POST")

	r.HandleFunc("/", service.GetHelloWorld).Methods("GET")

	r.HandleFunc("/ranking/{who}", service.GetRankingOfWhoService(wordDB)).Methods("GET")

	return r
}
