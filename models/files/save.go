package files

import (
	"database/sql"
	"mime/multipart"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

func SaveFile(fileload string, user string, file multipart.File) error {
	db, err := sql.Open("sqlite3", "file.db")
	if err != nil {
		return err
	}
	defer db.Close()

	// 保存文件到磁盘
	if _, err := os.Stat("/uploads"); os.IsNotExist(err) {
		os.Mkdir("/uploads", 0777)
	}
	filePath := "/uploads/" + fileload // 请根据实际路径修改
	out, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer out.Close()

	_, err = out.ReadFrom(file)
	if err != nil {
		return err
	}

	// 保存文件路径和用户信息到数据库
	stmt, err := db.Prepare("INSERT INTO files(path, user) VALUES(?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(filePath, user)
	if err != nil {
		return err
	}

	return nil
}
