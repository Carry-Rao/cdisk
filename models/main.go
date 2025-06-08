package models

import (
	"github.com/Carry-Rao/cdisk/models/files"
	"github.com/Carry-Rao/cdisk/models/system"
	"github.com/Carry-Rao/cdisk/models/users"
)

func Init() {
	users.InitDB()
	files.InitDB()
	system.InitDB()
}
