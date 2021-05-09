package routes

import (
	"cov-api/handlers"

	"github.com/gorilla/mux"
)

func GetNotifyRoutes(r *mux.Router) {
	r.HandleFunc("/notifyall", handlers.SendNotification).Methods("GET")
}
