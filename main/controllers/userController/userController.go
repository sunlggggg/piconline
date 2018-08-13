package userController

import (
	"net/http"
	"github.com/sunlggggg/piconline/main/service/userservice"
	"github.com/sunlggggg/piconline/main/entity"
	"time"
	"encoding/json"
)

// "username=sunlg&password=123456&email=sunlggggg@gmail.com" "http://127.0.0.1:8080/user"
func Register(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		username := r.PostFormValue("username")
		password := r.PostFormValue("password")
		email := r.PostFormValue("email")
		// 部分初始化
		statue, userId := userservice.Register(entity.User{Name: username, Email: email, Password: password, RegisterTime: uint64(time.Now().Unix())})

		res := map[string]interface{}{
			"status":  statue,
			"message": userId,
		}
		bytesRes, _ := json.Marshal(res)
		w.Write(bytesRes)
	}
}
