package main

import (
	"net/http"

	"github.com/Carry-Rao/cdisk/routes"
)

func main() {
	http.Handle("/", routes.Router())
	http.ListenAndServe(":8080", nil)
}
