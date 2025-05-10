package models

import (
	_ "github.com/mattn/go-sqlite3"
)

func Login(username string, password string) (bool, error) {
	result, err := datebase("user", "SELECT * FROM users WHERE username =? AND password =?", username, password)
}
