package filecontroller

import (
	"net/http"
	"github.com/sunlggggg/piconline/main/service/fileservice"
	"github.com/sunlggggg/piconline/main/code"
	"encoding/json"
	"github.com/sunlggggg/piconline/main/jwtutil"
)

// curl -d "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVfdGltZSI6MzYwMCwiY3JlYXRlX3RpbWUiOjE1MzQ2OTgxNzYsIm5hbWUiOiJzdW5sZyIsInVzZXJfaWQiOjQ5fQ==.7feb56f806934648dabae10a3b0afd10db645ff37921f063f3d4ce26fd4be0ca" "http://127.0.0.1:8080/fileroot"
func CreateRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		token := r.PostFormValue("token")
		jwt := new(jwtutil.JWT)
		if jwt.Decode(token) {
			userID := jwt.PayLoad.UserID
			fileRootId, err := fileservice.CreateRoot(userID)
			res := map[string]interface{}{
				"status":  code.FileRootCreateSuccess,
				"message": fileRootId,
			}
			if err != nil {
				res["status"] = code.FileRootCreateFail
			}
			bytesRes, _ := json.Marshal(res)
			w.Write(bytesRes)
		} else {
			res := map[string]interface{}{
				"status":  code.InValidToken,
				"message": "未授权请求",
			}
			bytesRes, _ := json.Marshal(res)
			w.Write(bytesRes)
		}
	}
}

// 上传文件
func UploadFile() {

}
