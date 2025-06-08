package models

import (
	"github.com/Carry-Rao/cdisk/models/system"
	users "github.com/Carry-Rao/cdisk/models/users"
)

func Init() {
	users.InitDB()
	system.InitDB()
}
