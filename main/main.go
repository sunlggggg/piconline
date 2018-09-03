package main

import (
	"net/http"
	"github.com/sunlggggg/piconline/main/controllers/pictureController"
	"github.com/sunlggggg/piconline/main/controllers/userController"
	"github.com/sunlggggg/piconline/main/config/mysqlconfig"
	"github.com/sunlggggg/piconline/main/controllers/filecontroller"
)

func main() {
	mysqlconfig.Init()

	// 文件
	// curl -d "token=eyJhbGc ... " "http://127.0.0.1:8080/root"
	http.HandleFunc("/root", filecontroller.CreateRoot)
	// curl -d "token=eyJhbGc ... &fatherID=13" "http://127.0.0.1:8080/dir"
	http.HandleFunc("/dir", filecontroller.CreateDir)
	// curl -F "token=eyJhbGc ... &fatherID=14&name=file1&isPublic=false&description=file01&file=@1.jpg" "http://127.0.0.1:8080/dir"
	http.HandleFunc("/file", filecontroller.UploadFile)
	// curl http://localhost:8080/picture -F "file=@1.jpg"
	http.HandleFunc("/picture", pictureController.Picture)

	// 用户
	// curl -d  "username=sunlg&password=123456&email=sunlggggg@gmail.com" "http://127.0.0.1:8080/user"
	http.HandleFunc("/user", userController.Register)

	// curl -d "username=sunlg&password=123456" "http://127.0.0.1:8080/login"
	http.HandleFunc("/login", userController.Login)

	http.ListenAndServe(":8080", nil)
}
