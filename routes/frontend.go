package routes

import (
	"github.com/Carry-Rao/cdisk/handlers"

	"github.com/gorilla/mux"
)

func frontendRoutes(r *mux.Router) {

}

func indexRoutes(r *mux.Router) error {
	return r.HandleFunc("/", handlers.indexHandler).Methods("GET")
}

func loginRoutes(r *mux.Router) error {
	return r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
}

func registerRoutes(r *mux.Router) error {
	return r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET", "POST")
}
