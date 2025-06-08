package models

func CreateUser(username string, password string, email string) error {
	_, err := db.Exec("INSERT INTO users (username, password, email) VALUES (?,?,?)", username, password, email)
	return err
}
