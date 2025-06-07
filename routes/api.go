package routes

import (
	"github.com/gorilla/mux"
)

func apiRoutes() *mux.Router {
}

func loginApiRoutes(r *mux.Router) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/login", handlers.loginApiHandler).Methods("POST")
	return r
}

func registerApiRoutes(r *mux.Router) *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/api/register", handlers.registerApiHandler).Methods("POST")
	return r
}
