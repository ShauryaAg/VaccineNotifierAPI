package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"cov-api/models/db"
	"cov-api/routes"
	"cov-api/utils"

	"github.com/rs/cors"
)

var (
	CURRENT_HOST = os.Getenv("CURRENT_HOST")
)

func main() {
	utils.AddCronJobs(CURRENT_HOST)

	db.DBCon, _ = db.CreateDatabase() // initialising the database

	r := routes.GetRoutes()

	// Serving Static files
	fs := http.FileServer(http.Dir("./static"))
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", fs))

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"POST", "GET", "PATCH", "OPTIONS"},
		AllowedHeaders:   []string{"Access-Control-Allow-Origin", "Accept", "Accept-Language", "Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})

	srv := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      c.Handler(r),
	}

	log.Fatal(srv.ListenAndServe())

}
