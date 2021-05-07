package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"cov-api/handlers"
	"cov-api/middlewares"
	"cov-api/models/db"
	"cov-api/utils"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

var (
	CURRENT_HOST = os.Getenv("CURRENT_HOST")
)

func main() {
	utils.AddCronJobs(CURRENT_HOST)

	db.DBCon, _ = db.CreateDatabase() // initialising the database

	r := mux.NewRouter().StrictSlash(true)

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "OPTIONS", "PUT"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Accept", "Accept-Language", "Content-Type"},
		AllowCredentials: true,
		Debug:            true,
	})

	r.HandleFunc("/t/{token}", handlers.VerifyToken).Methods("GET")      // GET /t/<token>; For Email verification
	r.HandleFunc("/u/{token}", handlers.UnsubscribeToken).Methods("GET") // GET /t/<token>; For Email verification

	// Prefix
	api := r.PathPrefix("/api").Subrouter()

	// API Routes
	api.HandleFunc("/login", handlers.Login).Methods("POST")       // POST /login
	api.HandleFunc("/register", handlers.Register).Methods("POST") // POST /register
	api.HandleFunc("/notifyall", handlers.SendNotification).Methods("GET")

	// Auth routes
	api.Handle("/user", middlewares.AuthMiddleware(
		http.HandlerFunc(handlers.GetUser),
	)).Methods("GET") // GET /user Auth: Bearer <Token>
	api.Handle("/user", middlewares.AuthMiddleware(
		http.HandlerFunc(handlers.UpdateUser),
	)).Methods("PATCH") // GET /user Auth: Bearer <Token>
	api.Handle("/unsubscribe", middlewares.AuthMiddleware(
		http.HandlerFunc(handlers.UnsubscribeUser),
	)).Methods("POST") // POST /unsubscribe Auth: Bearer <Token>

	srv := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      c.Handler(r),
	}

	log.Fatal(srv.ListenAndServe())

}
