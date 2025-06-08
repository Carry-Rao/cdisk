package routes

import (
	"net/http"

	"github.com/Carry-Rao/cdisk/handlers"

	"github.com/gorilla/mux"
)

func frontendRoutes(r *mux.Router) {
	indexRoutes(r)
	loginRoutes(r)
	registerRoutes(r)
	staticsRoutes(r)
}

func indexRoutes(r *mux.Router) {
	r.HandleFunc("/", handlers.IndexHandler).Methods("GET")
}

func loginRoutes(r *mux.Router) {
	r.HandleFunc("/login", handlers.LoginHandler).Methods("GET", "POST")
}

func registerRoutes(r *mux.Router) {
	r.HandleFunc("/register", handlers.RegisterHandler).Methods("GET", "POST")
}

func staticsRoutes(r *mux.Router) {
	r.PathPrefix("/statics/").Handler(http.StripPrefix("/statics/", http.FileServer(http.Dir("statics/"))))
}
