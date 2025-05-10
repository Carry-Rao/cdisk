package routes

import (
	"net/http"

	"github.com/Carry-Rao/cdisk/handlers"
)

func IndexHandler() {
	http.HandleFunc("/", handlers.Index)
}
