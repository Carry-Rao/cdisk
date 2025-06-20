package routes

import (
	"net/http"

	api "github.com/Carry-Rao/cdisk/handlers/api"
	"github.com/gorilla/mux"
)

func InitApi(router *mux.Router) {
	router.HandleFunc("/api/upload", api.ApiUploadHandler).Methods(http.MethodPut)
	// router.HandleFunc("/api/download", handlers.ApiDownloadHandler).Methods(http.MethodGet)
	// router.HandleFunc("/api/delete", handlers.ApiDeleteHandler).Methods(http.MethodDelete)
	// router.HandleFunc("/api/list", handlers.ApiListHandler).Methods(http.MethodGet)
	// router.HandleFunc("/api/info", handlers.ApiInfoHandler).Methods(http.MethodGet)
}
