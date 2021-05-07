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
	api.HandleFunc("/login", handlers.Login).Methods("POST")       // POST /login
	api.HandleFunc("/register", handlers.Register).Methods("POST") // POST /register

	// Auth routes
	api.Handle("/user", middlewares.AuthMiddleware(
		http.HandlerFunc(handlers.GetUser),
	)).Methods("GET") // GET /user Auth: Bearer <Token>
	api.Handle("/user", middlewares.AuthMiddleware(
		http.HandlerFunc(handlers.UpdateUser),
	)).Methods("PATCH") // GET /user Auth: Bearer <Token>
}
