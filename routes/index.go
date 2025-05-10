package routes

import (
	"net/http"

	"github.com/Carry-Rao/cdisk/handlers"
	"github.com/gorilla/mux"
)

func index(route *mux.Router) {
	route.HandleFunc("/", handlers.Index).Methods(http.MethodGet)
}
