package handlers

import (
	"net/http"

	users "github.com/Carry-Rao/cdisk/models/users"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	token := r.FormValue("token")
	if users.CheckWithToken(username, token) {
		http.ServeFile(w, r, "statics/index.html")
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
