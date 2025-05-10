package logger

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

var InfoDB *sql.DB

func InitInfoDB() {
	//检查是否存在logs文件夹，不存在则创建
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0777)
	}
	var err error
	InfoDB, err = sql.Open("sqlite3", "logs/info.db")
	if err != nil {
		fmt.Println("logger/InitInfoDB: ", err)
	}
	_, err = InfoDB.Exec("CREATE TABLE IF NOT EXISTS info (time TEXT, httperror TEXT, httpmethod TEXT, load TEXT, message TEXT)")
	if err != nil {
		fmt.Println("logger/InitInfoDB: ", err)
	}
}

func Info(httpCode int, httpMethod string, load string, message string) {
	//检测info.db格式是否正确，如果不正确则重新创建
	_, err := InfoDB.Exec("SELECT * FROM info")
	if err != nil {
		fmt.Println("logger/Info: ", err)
		InitInfoDB()
	}
	_, err = InfoDB.Exec("INSERT INTO info (time, httperror, httpmethod, load, message) VALUES (?, ?, ?, ?, ?)", time.Now(), httpCode, httpMethod, load, message)
	if err != nil {
		fmt.Println("logger/Info: ", err)
	}
}

func InitInfo() {
	//检测info.db是否存在，不存在则创建
	var err error
	if _, err = os.Stat("logs/info.db"); os.IsNotExist(err) {
		InitInfoDB()
	}
	InfoDB, err = sql.Open("sqlite3", "logs/info.db")
	if err != nil {
		fmt.Println("logger/InitInfo: ", err)
	}
}

func CloseInfo() {
	InfoDB.Close()
}

func GetInfo() []map[string]interface{} {
	rows, err := InfoDB.Query("SELECT * FROM info")
	if err != nil {
		fmt.Println("logger/GetInfo: ", err)
	}
	defer rows.Close()
	result := []map[string]interface{}{}
	for rows.Next() {
		var time string
		var httperror int
		var httpmethod string
		var load string
		var message string
		err = rows.Scan(&time, &httperror, &httpmethod, &load, &message)
		if err != nil {
			fmt.Println("logger/GetInfo: ", err)
		}
		result = append(result, map[string]interface{}{"time": time, "httperror": httperror, "httpmethod": httpmethod, "load": load, "message": message})
	}
	return result
}
