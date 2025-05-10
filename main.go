package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/Carry-Rao/cdisk/logger"
	"github.com/Carry-Rao/cdisk/routes"
	_ "github.com/mattn/go-sqlite3"
	"gopkg.in/ini.v1"
)

var config *ini.File

func initDB() {
	if _, err := os.Stat("logs"); os.IsNotExist(err) {
		os.Mkdir("logs", 0777)
	}
	var err error
	if _, err = os.Stat("logs/error.db"); os.IsNotExist(err) {
		ErrorDB, _ := sql.Open("sqlite3", "logs/error.db")
		_, err = ErrorDB.Exec("CREATE TABLE errors (TEXT, httperror TEXT, httpmethod TEXT, load TEXT, message TEXT)")
		if err != nil {
			fmt.Println("logger/InitErrorDB: ", err)
		}
	}
	if _, err = os.Stat("logs/info.db"); os.IsNotExist(err) {
		InfoDB, _ := sql.Open("sqlite3", "logs/info.db")
		_, err = InfoDB.Exec("CREATE TABLE info (TEXT, httperror TEXT, httpmethod TEXT, load TEXT, message TEXT)")
		if err != nil {
			fmt.Println("logger/InitInfoDB: ", err)
		}
	}
}

func output() {
	fmt.Print(`   ____       _   _         _    
  / ___|   __| | (_)  ___  | | __
 | |      / _` + "`" + ` | | | / __| | |/ /
 | |___  | (_| | | | \__ \ |   < 
  \____|  \__,_| |_| |___/ |_|\_\` + ``)
	fmt.Println("\n-----------------------------------")
	fmt.Println("v0.1.0 - A simple web disk system")
}

func main() {
	output()
	initDB()
	var err error
	config, err = ini.Load("config.ini")
	if err != nil {
		logger.Error(0, "/nothttp", "nothttp", "Failed to load config file: "+err.Error())
		return
	}
	port := config.Section("server").Key("port").String()
	logger.Info(0, "/nothttp", "nothttp", "Server started on port "+port)
	http.Handle("/", routes.Router())
	http.ListenAndServe(port, nil)
}
