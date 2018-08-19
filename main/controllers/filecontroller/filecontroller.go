package filecontroller

import (
	"net/http"
	"github.com/sunlggggg/piconline/main/service/fileservice"
	"github.com/sunlggggg/piconline/main/code"
	"encoding/json"
)
// curl -d "username=sunlg" "http://127.0.0.1:8080/fileroot"
func CreateRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		// TODO token 权限
		name := r.PostFormValue("username")
		fileRootId, err := fileservice.CreateRoot(name)
		res := map[string]interface{}{
			"status":  code.FileRootCreateSuccess,
			"message": fileRootId,
		}
		if err != nil {
			res["status"] = code.FileRootCreateFail
		}
		bytesRes, _ := json.Marshal(res)
		w.Write(bytesRes)
	}
}

// 上传文件
func UploadFile() {

}