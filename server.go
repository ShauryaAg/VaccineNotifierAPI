package main

import (
	"log"
	"net/http"
	"time"

	"cov-api/handlers"
	"cov-api/middlewares"
	"cov-api/models/db"

	"github.com/gorilla/mux"
)

func main() {
	db.DBCon, _ = db.CreateDatabase() // initialising the database

	r := mux.NewRouter().StrictSlash(true)

	// API Routes
	r.HandleFunc("/login", handlers.Login).Methods("POST")       // POST /login
	r.HandleFunc("/register", handlers.Register).Methods("POST") // POST /register

	r.Handle("/user", middlewares.AuthMiddleware(
		http.HandlerFunc(handlers.GetUser),
	)).Methods("GET") // GET /user Auth: Bearer <Token>

	srv := &http.Server{
		Addr:         ":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      r,
	}

	log.Fatal(srv.ListenAndServe())

}
