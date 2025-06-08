package routes

import (
	"github.com/Carry-Rao/cdisk/handlers"
	"github.com/gorilla/mux"
)

func loginApiRoutes(r *mux.Router) {
	r.HandleFunc("/api/users/login", handlers.LoginApiHandler).Methods("POST")
}

func registerApiRoutes(r *mux.Router) {
	r.HandleFunc("/api/users/register", handlers.RegisterApiHandler).Methods("POST")
}

func get_name(r *mux.Router) {
	r.HandleFunc("/api/system/get_name", handlers.GetName).Methods("GET")
}

func apiRoutes(r *mux.Router) {
	loginApiRoutes(r)
	registerApiRoutes(r)
	get_name(r)
}
