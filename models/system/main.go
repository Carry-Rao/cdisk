package system

import (
	"database/sql"

	"github.com/Carry-Rao/cdisk/log"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func InitDB() {
	var err error
	db, err = sql.Open("sqlite3", "db/system.db")
	if err != nil {
		log.Error("models/system/InitDB()", err.Error())
		return
	}
	db.Exec(`
		CREATE TABLE IF NOT EXISTS system (
			name TEXT
		);
	`)
	rows, err := db.Query("SELECT name FROM system")
	if err != nil {
		log.Error("models/system/InitDB()", err.Error())
		return
	}
	defer rows.Close()
	if !rows.Next() {
		db.Exec("INSERT INTO system (name) VALUES ('Cdisk网盘')")
	}
}
