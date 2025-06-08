package log

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var db_error *sql.DB

func initErrorDB() {
	var err error
	if _, err := os.Stat("db/logs"); os.IsNotExist(err) {
		os.Mkdir("db/logs", 0777)
	}
	db_error, err = sql.Open("sqlite3", "db/logs/error.db")
	if err != nil {
		panic(err)
	}
	db_error.Exec(`
		CREATE TABLE IF NOT EXISTS errors (
			time TEXT,
			func TEXT,
			error TEXT
		)
	`)
	db_error.Exec(`
		CREATE TABLE IF NOT EXISTS http_errors (
			time TEXT,
			route TEXT,
			error TEXT,
			status INT
		)
	`)
}

func Error(f string, Err string) {
	time := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db_error.Prepare("INSERT INTO errors (time, func, error) VALUES (?,?,?)")
	if err != nil {
		panic(err)
	}
	stmt.Exec(time, f, Err)
}

func closeErrorDB() {
	db_error.Close()
}

func HttpError(route string, Err string, statue int) {
	time := time.Now().Format("2006-01-02 15:04:05")
	stmt, err := db_error.Prepare("INSERT INTO http_errors (time, route, error, status) VALUES (?,?,?,?)")
	if err != nil {
		panic(err)
	}
	stmt.Exec(time, route, Err, statue)
}
