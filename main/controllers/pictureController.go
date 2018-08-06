package controllers

import (
	"net/http"
	"fmt"
	"strings"
	"github.com/sunlggggg/piconline/code"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()     //解析参数，默认是不会解析的
	fmt.Println(r.Form) //这些信息是输出到服务器端的打印信息
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	fmt.Fprintf(w, "Hello astaxie!")
}

func Picture(w http.ResponseWriter, r *http.Request){
	// 读取请求类型
	switch r.Method {
	case http.MethodGet:
		println("rest get ...  ")
	    // 获取图片
	case http.MethodPost:
		// 上传图片
		println("rest post ... ")
	
	case http.MethodDelete:
		println("this is delete ")
	case http.MethodPut:
		println("this is put ")
	default:
		print("error code",code.INVAILD_REQUEST)
	}
}
