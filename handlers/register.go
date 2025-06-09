package handlers

import (
	"net/http"

	model "github.com/Carry-Rao/cdisk/models/users"
)

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "statics/register.html")
}

func RegisterApiHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	username := r.Form.Get("username")
	password := r.Form.Get("password")
	email := r.Form.Get("email")
	if model.CheckUserExists(username) {
		http.Redirect(w, r, "/register", http.StatusFound)
		w.Write([]byte("Username already exists"))
		return
	}
	if model.CreateUser(username, password, email) != nil {
		http.Redirect(w, r, "/register", http.StatusFound)
		w.Write([]byte("Error creating user"))
	} else {
		http.Redirect(w, r, "/login", http.StatusFound)
		w.Write([]byte("User created successfully"))
	}
}
