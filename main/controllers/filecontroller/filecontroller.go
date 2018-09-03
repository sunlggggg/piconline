package filecontroller

import (
	"net/http"
	"github.com/sunlggggg/piconline/main/service/fileservice"
	"github.com/sunlggggg/piconline/main/code"
	"encoding/json"
	"github.com/sunlggggg/piconline/main/jwtutil"
	"fmt"
	"strconv"
	"log"
)

// curl -d "token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVfdGltZSI6MzYwMCwiY3JlYXRlX3RpbWUiOjE1MzQ2OTgxNzYsIm5hbWUiOiJzdW5sZyIsInVzZXJfaWQiOjQ5fQ==.7feb56f806934648dabae10a3b0afd10db645ff37921f063f3d4ce26fd4be0ca" "http://127.0.0.1:8080/fileroot"
func CreateRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		token := r.PostFormValue("token")
		jwt := new(jwtutil.JWT)
		if jwt.Decode(token) {
			userID := jwt.PayLoad.UserID
			fileRootId, err := fileservice.CreateRootDir(userID)
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

// curl -d "token=eyJhbGc ... &fatherID=13" "http://127.0.0.1:8080/dir"
func CreateDir(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		token := r.PostFormValue("token")
		jwt := new(jwtutil.JWT)
		if jwt.Decode(token) {
			userID := jwt.PayLoad.UserID
			strFatherId := r.PostFormValue("fatherID")
			fatherId, err := strconv.ParseUint(strFatherId, 10, 64)
			res := map[string]interface{}{
				"status":  nil,
				"message": nil,
			}

			if err != nil {
				log.Println(err)
				res["status"] = code.InValidArgs
			} else {
				dirId, err := fileservice.CreateDir(userID, fatherId)
				if err != nil {
					res["status"] = code.FileDirCreateFail
				} else {
					res["status"] = code.FileDirCreateSuccess
					res["message"] = dirId
				}
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
// curl -F "file=@1.jpg" "http://127.0.0.1:8080/file?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHBpcmVfdGltZSI6MzYwMCwiY3JlYXRlX3RpbWUiOjE1MzU5NjQ5NDAsIm5hbWUiOiJzdW5sZyIsInVzZXJfaWQiOjQ5fQ==.a34f52db47784547b65266578527e0d5dc6990e416e850a2e118afe6a530079e&fatherID=14&name=file1&isPublic=0&description=file01"
func UploadFile(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		token := r.URL.Query().Get("token")
		jwt := new(jwtutil.JWT)
		if jwt.Decode(token) {
			userID := jwt.PayLoad.UserID
			fmt.Fprintln(w, "upload ok!")
			strFatherId := r.URL.Query().Get("fatherID")
			fatherId, err := strconv.ParseUint(strFatherId, 10, 64)
			name := r.URL.Query().Get("name")
			strIsPublic := r.URL.Query().Get("isPublic")
			isPublic, err := strconv.ParseBool(strIsPublic)
			description := r.URL.Query().Get("description")
			res := map[string]interface{}{
				"status":  nil,
				"message": nil,
			}
			if err != nil {
				log.Println(err)
				res["status"] = code.InValidArgs
			} else {
				fileID, err := fileservice.UploadFile(userID, fatherId, name, isPublic, description, r)
				if err != nil {
					res["status"] = code.FileCreateFail
					res["message"] = err
				} else {
					res["status"] = code.FileCreateSuccess
					res["message"] = fileID
				}
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
