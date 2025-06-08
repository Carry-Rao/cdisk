package models

import (
	"github.com/Carry-Rao/cdisk/models/file"
	"github.com/Carry-Rao/cdisk/models/system"
	"github.com/Carry-Rao/cdisk/models/user"
)

func Init() {
	user.InitDB()
	file.InitDB()
	system.InitDB()
}
