package routes

import (
	"cov-api/handlers"

	"github.com/gorilla/mux"
)

func GetTokenRoutes(r *mux.Router) {
	r.HandleFunc("/t/{token}", handlers.VerifyToken).Methods("GET")      // GET /t/<token>; For Email verification
	r.HandleFunc("/u/{token}", handlers.UnsubscribeToken).Methods("GET") // GET /t/<token>; For Email verification
}
