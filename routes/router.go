package routes

import (
	"github.com/gorilla/mux"
	"naratmalsami-survey-server/db"
	"naratmalsami-survey-server/service"
)

func SetupRouter(wordDB *db.WordDB) *mux.Router {
	r := mux.NewRouter()

	r.HandleFunc("/sheet", service.GetSheetService(wordDB)).Methods("GET")

	r.HandleFunc("/", service.GetHelloWorld).Methods("GET")

	return r
}
