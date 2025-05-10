package logger

import (
	"database/sql"
	"fmt"
	"os"
	"runtime"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

func Error(format string, a ...interface{}) {
	_, file, line, _ := runtime.Caller(1)
	fmt.Fprintf(os.Stderr, "Error at %s:%d: "+format+"\n", append([]interface{}{file, line}, a...)...)
}
