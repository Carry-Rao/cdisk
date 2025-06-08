package models

func CheckWithUAP(username string, password string) bool {
	return true
}

func CheckUserExists(username string) bool {
	db.Prepare("SELECT * FROM users WHERE username =?")
	rows, err := db.Query("SELECT * FROM users WHERE username =?", username)
	if err != nil {
		return true
	}
	defer rows.Close()
	if rows.Next() {
		return true
	}
	return false
}
