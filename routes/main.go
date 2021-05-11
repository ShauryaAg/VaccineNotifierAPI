package routes

import "github.com/gorilla/mux"

func GetRoutes() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)

	// Prefix
	api := r.PathPrefix("/api").Subrouter()

	GetTokenRoutes(r) // has no prefix 'api' in front of it

	// with prefix 'api'
	GetUserRoutes(api)
	GetNotifyRoutes(api)

	return r
}
