package models

import (
	"encoding/json"
	"math/rand"
)

func CheckWithUAP(username string, password string) bool {
	rows, err := db.Query("SELECT password FROM users WHERE username =?", username)
	if err != nil {
		return false
	}
	defer rows.Close()
	if rows.Next() {
		var Password string
		err := rows.Scan(&Password)
		if err != nil {
			return false
		}
		if password == Password {
			return true
		}
	}
	return false
}

func CheckUserExists(username string) bool {
	db.Prepare("SELECT * FROM users WHERE username =?")
	rows, err := db.Query("SELECT * FROM users WHERE username =?", username)
	if err != nil {
		return false
	}
	defer rows.Close()
	if rows.Next() {
		return true
	}
	return false
}

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func newToken() string {
	b := make([]byte, 32)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func CheckWithToken(username string, token string) bool {
	if !CheckUserExists(username) {
		return false
	}
	rows, err := db.Query("SELECT tokens FROM users WHERE username =?", username)
	if err != nil {
		return false
	}
	defer rows.Close()
	if rows.Next() {
		var tokens []string
		err := rows.Scan(&tokens)
		if err != nil {
			return false
		}
		for _, t := range tokens {
			if t == token {
				return true
			}
		}
	}
	return false
}

func GetToken(username string) string {
	token := newToken()
	rows, err := db.Query("SELECT tokens FROM users WHERE username =?", username)
	if err != nil {
		return ""
	}
	defer rows.Close()
	if rows.Next() {
		var tokens []string
		err := rows.Scan(&tokens)
		if err != nil {
			return ""
		}
		tokens = append(tokens, token)
		jsonTokens, err := json.Marshal(tokens)
		if err != nil {
			return ""
		}
		_, err = db.Exec("UPDATE users SET tokens =? WHERE username =?", string(jsonTokens), username)
		if err != nil {
			return ""
		}
		return token
	}
	return token
}
