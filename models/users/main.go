package models

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "db/users.db")
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			uid integer,
			username text,
			password text,
			email text,
			gid integer,
			token json
		);
	`)
	if err != nil {
		panic(err)
	}
}
