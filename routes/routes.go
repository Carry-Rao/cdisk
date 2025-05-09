package routes

import (
	"net/http"

	"github.com/gorilla/mux"
)

func InitRoutes() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", IndexHandler).Methods("GET")
	router.HandleFunc("/upload", UploadHandler).Methods("POST")
	router.HandleFunc("/download/{file_name}", DownloadHandler).Methods("GET")
	router.HandleFunc("/delete/{file_name}", DeleteHandler).Methods("GET")

	return router
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
}
