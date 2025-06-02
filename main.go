package main

import (
	"fmt"

	"github.com/Carry-Rao/cdisk/routes"

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
	fmt.Printf("This is a webdisk system by carry.")
}

func initAll(r *mux.Router) error {
	err := routes.InitRoutes(r)
	if err != nil {
		return err
	}
}

func main() {
	hello()
	var r *mux.Router = mux.NewRouter()
	err := initAll(r)
}
