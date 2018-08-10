package main

import (
	"github.com/sunlggggg/piconline/main/controllers"
	"github.com/sunlggggg/piconline/main/utils"
	"net/http"
)

type Rect struct {
	width, length float64
}

func (rect Rect) area() float64 {
	return rect.width * rect.length
}

func main() {
	name := "main/data/1.txt"
	utils.Write(name, "dd")
	http.HandleFunc("/picture", controllers.Picture)
	http.ListenAndServe(":8080", nil)
}
