package main

import (
	"github.com/sunlggggg/piconline/main/controllers"
	"net/http"
	"github.com/sunlggggg/piconline/main/utils"
)

func main() {
	name := "main/data/1.txt"
	utils.Write(name, "dd")
	http.HandleFunc("/picture", controllers.Picture)
	http.ListenAndServe(":8080", nil)
}