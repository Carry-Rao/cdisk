package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

func Login(username string, password string) (bool, error) {
	db, err := sql.Open("sqlite3", "database/users.db")
	if err != nil {
		return false, err
	}
	defer db.Close()
	var result string
	err = db.QueryRow("SELECT password FROM users WHERE username = ?", username).Scan(&result)
	if err != nil {
		return false, err
	}
	if result == password {
		return true, nil
	} else {
		return false, nil
	}
}

func AuthToken(token string) (string, error) {
	db, err := sql.Open("sqlite3", "database/users.db")
	if err != nil {
		return "", err
	}
	defer db.Close()
	var result string
	err = db.QueryRow("SELECT username FROM users WHERE token = ?", token).Scan(&result)
	if err != nil {
		return "", err
	}
	if result == token {
		return result, nil
	} else {
		return "", nil
	}
}
