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
		w.Write([]byte("User already exists"))
		return
	}
	if model.CreateUser(username, password, email) != nil {
		w.Write([]byte("User created successfully"))
	} else {
		w.Write([]byte("Error creating user"))
	}
}
