package main

import (
	"net/http"
	"github.com/sunlggggg/piconline/main/controllers/pictureController"
	"github.com/sunlggggg/piconline/main/controllers/userController"
	"github.com/sunlggggg/piconline/main/config/mysql"
	"github.com/sunlggggg/piconline/main/controllers/filecontroller"
)

func main() {
	mysql.Init()
	http.HandleFunc("/fileRoot", filecontroller.CreateRoot)
	http.HandleFunc("/picture", pictureController.Picture)
	http.HandleFunc("/user", userController.Register)
	http.ListenAndServe(":80", nil)
}
