package main

import (
	"net/http"
	"github.com/sunlggggg/piconline/controllers"
)


func main() {
	http.HandleFunc("/", controllers.Hello)
	http.ListenAndServe(":8080",nil)
}