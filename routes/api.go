package routes

import (
	"net/http"

	"github.com/Carry-Rao/webdisk/handlers"
	"github.com/gorilla/mux"
)

func api(router *mux.Router) {
	http.HandleFunc("/api/upload", handlers.UploadHandler)
	http.HandleFunc("/api/download", handlers.DownloadHandler)
	http.HandleFunc("/api/delete", handlers.DeleteHandler)
	http.HandleFunc("/api/list", handlers.ListHandler)
}
