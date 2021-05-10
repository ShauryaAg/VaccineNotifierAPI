package routes

import (
	"net/http"

	"cov-api/handlers"
	"cov-api/middlewares"

	"github.com/gorilla/mux"
)

func GetUserRoutes(r *mux.Router) {
	// Prefix
	api := r.PathPrefix("/auth").Subrouter()

	// API Routes
	api.HandleFunc("/login", handlers.Login).Methods("POST")                      // POST /api/auth/login
	api.HandleFunc("/register", handlers.Register).Methods("POST")                // POST /api/auth/register
	api.HandleFunc("/reset_password", handlers.ResetUserPassword).Methods("POST") // POST /api/auth/reset_password

	// Auth routes
	api.Handle("/user", middlewares.AuthMiddleware(
		http.HandlerFunc(handlers.GetUser),
	)).Methods("GET") // GET /api/auth/user Auth: Bearer <Token>
	api.Handle("/user", middlewares.AuthMiddleware(
		http.HandlerFunc(handlers.UpdateUser),
	)).Methods("PATCH") // PATCH /api/auth/user Auth: Bearer <Token>
	api.Handle("/unsub", middlewares.AuthMiddleware(
		http.HandlerFunc(handlers.UnsubscribeUser),
	)).Methods("POST") // POST /api/auth/unsub Auth: Bearer <Token>

}
