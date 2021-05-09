package routes

import (
	"cov-api/handlers"
	"cov-api/views"

	"github.com/gorilla/mux"
)

func GetTokenRoutes(r *mux.Router) {
	r.HandleFunc("/t/{token}", handlers.VerifyToken).Methods("GET")      // GET /t/<token>; For Email verification
	r.HandleFunc("/u/{token}", handlers.UnsubscribeToken).Methods("GET") // GET /u/<token>; For Unsubscribing to Emails

	// views
	r.HandleFunc("/f/{token}", views.ResetPasswordView).Methods("GET") // GET /f/<token>; For Password Reset Emails
	r.HandleFunc("/f/{token}", views.ResetPassword).Methods("POST")    // POST /f/<token>; For Password Reset Emails
}
