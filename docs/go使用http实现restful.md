# 使用go实现rest
``` go
// post : curl http://localhost:8080/picture -F "pic=@1.jpg"
func Picture(w http.ResponseWriter, r *http.Request){
	// 读取请求类型
	switch r.Method {
	case http.MethodGet:
		// 下载图片
		println("rest get ...  ")
	    fout ,err:=os.Open("upload/1.jpg")
	    if err != nil {
	    	println(err)
		}
		defer fout.Close()
	    buf := make([]byte,1024)
	    for {
	    	n, _:= fout.Read(buf)
	    	if 0 == n {
	    		break
			}
			println(n)
			w.Write(buf[:n])
		}
	case http.MethodPost:
		// 上传图片
		println("rest post ... ")
		r.ParseMultipartForm(32 << 20)
		file, handler, err := r.FormFile("pic")
		if err != nil {
			fmt.Println(err)
			return
		}
		defer file.Close()
		//0666 赋予读写权限
		f, err := os.OpenFile("upload/" + handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
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
		print("error code",code.InvalidRequest)
	}
}
```