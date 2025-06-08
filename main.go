package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/Carry-Rao/cdisk/log"
	"github.com/Carry-Rao/cdisk/models"
	"github.com/Carry-Rao/cdisk/routes"
	"github.com/go-ini/ini"
	"github.com/gorilla/mux"
)

func hello() {
	fmt.Printf(`
   ____       _   _         _    
  / ___|   __| | (_)  ___  | | __
 | |      / _ ` + "`" + `| | | / __| | |/ /
 | |___  | (_| | | | \__ \ |   < 
  \____|  \__,_| |_| |___/ |_|\_\
                                                       
`)
	fmt.Println("This is a webdisk system by carry.")
}

func initAll(r *mux.Router) {
	routes.InitRoutes(r)
	//检查是否有db文件夹，没有则创建
	if _, err := os.Stat("db"); os.IsNotExist(err) {
		os.Mkdir("db", 0777)
	}
	log.InitDB()
	models.Init()
}

func main() {
	hello()
	var r *mux.Router = mux.NewRouter()
	initAll(r)
	config, err := ini.Load("config.ini")
	if err != nil {
		fmt.Println("config.ini not found")
	}
	port := config.Section("server").Key("port").String()
	fmt.Println("Server running on port ", port)
	http.ListenAndServe(port, r)
}
