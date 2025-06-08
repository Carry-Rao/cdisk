package handlers

import (
	"net/http"

	model "github.com/Carry-Rao/cdisk/models/system"
)

func GetName(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(model.GetName()))
}
