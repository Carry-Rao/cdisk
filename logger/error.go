package logger

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var ErrorDB *sql.DB

func InitErrorDB() {
	//检查是否存在logs文件夹，不存在则创建

}

func Error(httpCode int, httpMethod string, load string, message string) {
	_, err := ErrorDB.Exec("INSERT INTO errors (time, httperror, httpmethod, load, message) VALUES (?, ?, ?, ?, ?)", time.Now(), httpCode, httpMethod, load, message)
	if err != nil {
		fmt.Println("logger/Error: ", err)
	}
}

func GetErrors() []map[string]interface{} {
	rows, err := ErrorDB.Query("SELECT * FROM errors")
	if err != nil {
		fmt.Println("logger/GetErrors: ", err)
	}
	defer rows.Close()
	errors := []map[string]interface{}{}
	for rows.Next() {
		var time string
		var httperror int
		var httpmethod string
		var load string
		var message string
		err := rows.Scan(&time, &httperror, &httpmethod, &load, &message)
		if err != nil {
			fmt.Println("logger/GetErrors: ", err)
		}
		errors = append(errors, map[string]interface{}{"time": time, "httperror": httperror, "httpmethod": httpmethod, "load": load, "message": message})
	}
	return errors
}

func InitError() {
	var err error
	ErrorDB, err = sql.Open("sqlite3", "logs/error.db")
	if err != nil {
		fmt.Println("logger/InitError: ", err)
	}
}

func CloseError() {
	ErrorDB.Close()
}
