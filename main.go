package main

import (
	"fmt"
	"net/http"

	"github.com/Carry-Rao/cdisk/logger"
	"github.com/Carry-Rao/cdisk/routes"
	"gopkg.in/ini.v1"
)

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
	config, err := ini.Load("config.ini")
	if err != nil {
		logger.Error(0, "/nothttp", "nothttp", "Failed to load config file: "+err.Error())
		return
	}
	port := config.Section("server").Key("port").String()
	logger.Info(0, "/nothttp", "nothttp", "Server started on port "+port)
	http.Handle("/", routes.Router())
	http.ListenAndServe(port, nil)
}
