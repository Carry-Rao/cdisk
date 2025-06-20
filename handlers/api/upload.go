package handlers

import (
	"net/http"

	files "github.com/Carry-Rao/cdisk/models/files"
	auth "github.com/Carry-Rao/cdisk/models/login"
)

func ApiUploadHandler(w http.ResponseWriter, r *http.Request) {
	//获取上传的file文件
	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	defer file.Close()
	//获取文件名
	filename := header.Filename
	//获取cookie中的token
	token := r.Header.Get("Authorization")
	//验证token
	if token == "" {
		http.Error(w, "token is empty", http.StatusUnauthorized)
		return
	}
	user, err := auth.AuthToken(token)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	//保存文件到服务器
	err = files.SaveFile(filename, user, file)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	//返回上传成功的响应
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte("{\"code\": 0, \"msg\": \"upload success\"}"))
}
