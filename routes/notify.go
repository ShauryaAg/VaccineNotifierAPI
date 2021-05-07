package routes

import (
	"cov-api/handlers"

	"github.com/gorilla/mux"
)

func GetNotifyRoutes(r *mux.Router) {
	// Prefix
	api := r.PathPrefix("/token").Subrouter()

	api.HandleFunc("/notifyall", handlers.SendNotification).Methods("GET")
}
