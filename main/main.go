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
	// curl
	http.HandleFunc("/fileroot", filecontroller.CreateRoot)
	http.HandleFunc("/picture", pictureController.Picture)
	http.HandleFunc("/user", userController.Register)
	http.HandleFunc("/login", userController.Login)
	http.ListenAndServe(":8080", nil)
}


