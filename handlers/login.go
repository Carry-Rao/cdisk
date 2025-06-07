package handlers

import (
	"net/http"

	model "github.com/Carry-Rao/webdisk/models/users"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/login.html")
}

func LoginApiHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.FormValue("username")
	password := r.FormValue("password")
	user := model.User{Username: username, Password: password}
	if model.CheckWithUAP(user)() {
		http.Redirect(w, r, "/", http.StatusFound)
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
	}
}
