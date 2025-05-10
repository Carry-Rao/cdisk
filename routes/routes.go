package routes

import (
	"net/http"

	"github.com/Carry-Rao/cdisk/handlers"
	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handlers.IndexHandler).Methods("GET")
	router.HandleFunc("/upload", handlers.UploadHandler).Methods("POST")
	router.HandleFunc("/download/{file_name}", handlers.DownloadHandler).Methods("GET")
	router.HandleFunc("/delete/{file_name}", handlers.DeleteHandler).Methods("GET")

	return router
}

func Router() http.Handler {
	return InitRoutes()
}
