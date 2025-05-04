package routes

import (
	"net/http"

	"github.com/Carry-Rao/webdisk/handlers"
	"github.com/Carry-Rao/webdisk/middlewares"
	"github.com/Carry-Rao/webdisk/utils"
	"github.com/gorilla/mux"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter()

	r.Use(middlewares.LoggerMiddleware)
	r.Use(middlewares.RecoveryMiddleware)

	r.HandleFunc("/api/v1/upload", handlers.UploadHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/download", handlers.DownloadHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/delete", handlers.DeleteHandler).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/list", handlers.ListHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/info", handlers.InfoHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/share", handlers.ShareHandler).Methods(http.MethodPost)
	r.HandleFunc("/api/v1/unshare", handlers.UnshareHandler).Methods(http.MethodDelete)
	r.HandleFunc("/api/v1/share/list", handlers.ShareListHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/share/info", handlers.ShareInfoHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/share/download", handlers.ShareDownloadHandler).Methods(http.MethodGet)
	r.HandleFunc("/api/v1/share/delete", handlers.ShareDeleteHandler).Methods(http.MethodDelete)

	r.PathPrefix("/").Handler(http.StripPrefix("/", http.FileServer(http.Dir(utils.GetAbsPath("web/dist/")))))

	return r
}
