package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	index(router)
	InitApi(router)
	return router
}

func Router() http.Handler {
	return InitRoutes()
}
