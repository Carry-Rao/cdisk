package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Carry-Rao/cdisk/logger"
	_ "github.com/mattn/go-sqlite3"
)

func datebase(sqlName string, command string, args ...interface{}) (sql.Result, error) {
	if sqlName == "" {
		logger.Error(500, "", "", "Database not initialized")
		return nil, fmt.Errorf("Database not initialized")
	}
	if sqlName == "user" {
		return dbUser(command, args...)
	}
	return nil, fmt.Errorf("Invalid database")
}

func dbUser(command string, args ...interface{}) (sql.Result, error) {
	//检测数据库是否存在，不存在则创建
	if _, err := os.Stat("user.db"); os.IsNotExist(err) {
		file, err := os.Create("user.db")
		if err != nil {
			fmt.Println(err)
		}
		defer file.Close()
		var sqlcode = `CREATE TABLE users (
			id INTEGER,
			username VARCHAR(255) NOT NULL,
			password VARCHAR(255) NOT NULL,
			email VARCHAR(255) NOT NULL,
			group INTEGER NOT NULL,
			userstatus INTEGER NOT NULL,
			capacity INTEGER NOT NULL,
		);`
		db, err := sql.Open("sqlite3", "user.db")
		_, err = db.Exec(sqlcode)
		if err != nil {
			fmt.Println(err)
		}
	}
	db, err := sql.Open("sqlite3", "user.db")
	if err != nil {
		logger.Error(500, "", "", "Database not initialized")
		os.Exit(1)
	}
	defer db.Close()
	return db.Exec(command, args...)
}
