package pictureController

import (
	"fmt"
	"github.com/sunlggggg/piconline/main/code"
	"io"
	"net/http"
	"os"
	"strings"
	"github.com/sunlggggg/piconline/main/utils"
	"time"
)

const (
	savePath = "/Users/sunligang/go-workspace/piconline/upload/"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()       //解析参数，默认是不会解析的
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

// post : curl http://localhost:8080/picture -F "file=@1.jpg"
func Picture(w http.ResponseWriter, r *http.Request) {
	// 读取请求类型
	switch r.Method {
	case http.MethodGet:
		// 下载图片
		println("rest get ...  ")
		fout, err := os.Open(savePath + "/1.jpg")
		if err != nil {
			println(err)
		}
		defer fout.Close()
		buf := make([]byte, 1024)
		for {
			n, _ := fout.Read(buf)
			if 0 == n {
				break
			}
			println(n)
			w.Write(buf[:n])
		}
	case http.MethodPost:
		// 上传文件
		println("rest post ... ")
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("file")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//0666 赋予读写权限
		filename := time.ANSIC + utils.Hash(handler.Filename, 10)
		println(filename)
		f, err := os.OpenFile(savePath+"1.jpg", os.O_WRONLY|os.O_CREATE, 0666)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()
		io.Copy(f, file)
		fmt.Fprintln(w, "upload ok!")
	case http.MethodDelete:
		println("this is delete ")
	case http.MethodPut:
		println("this is put ")
	default:
		print("error code", code.InvalidRequest)
	}
}
