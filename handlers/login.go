package handlers

import (
	"net/http"

	model "github.com/Carry-Rao/cdisk/models/users"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "statics/login.html")
}

func LoginApiHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	if model.CheckWithUAP(username, password) {
		http.SetCookie(w, &http.Cookie{
			Name:  "username",
			Value: username,
			Path:  "/",
		})
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
