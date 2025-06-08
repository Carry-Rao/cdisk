package routes

import "github.com/gorilla/mux"

func InitRoutes(r *mux.Router) {
	frontendRoutes(r)
	apiRoutes(r)
}
