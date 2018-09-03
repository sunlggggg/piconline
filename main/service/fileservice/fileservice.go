package fileservice

import (
	"github.com/sunlggggg/piconline/main/dao/filedao"
	"time"
	"github.com/sunlggggg/piconline/main/config/mysqlconfig"
	"net/http"
	"fmt"
	"github.com/sunlggggg/piconline/main/utils"
	"os"
	"io"
	"github.com/sunlggggg/piconline/main/dao/contentdao"
)

var fileUploadUrl = "file"

func CreateRootDir(userId uint64) (int64, error) {
	return filedao.InsertRootDir(userId, uint64(time.Now().Unix()))
}
func CreateDir(userId uint64, fatherId uint64) (int64, error) {
	return filedao.InsertDir(userId, fatherId, uint64(time.Now().Unix()))
}

func UploadFile(userId uint64, fatherId uint64, name string, isPublic bool, description string, r *http.Request) (int64, error) {
	// 事务测试 ...
	// 注意 在使用中应该使用tx而不是db
	db := mysqlconfig.Mysqldb
	tx, _ := db.Begin()
	// 文件上传 ...
	r.ParseMultipartForm(32 << 20)
	file, handler, err := r.FormFile(fileUploadUrl)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	defer file.Close()
	//0666 赋予读写权限
	filename := time.ANSIC + utils.Hash(handler.Filename, 10)
	f, err := os.OpenFile("upload/"+filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return -1, err
	}
	defer f.Close()
	io.Copy(f, file)

	nowTime := uint64(time.Now().Unix())
	contentId, err := contentdao.Insert(tx, name, isPublic, nowTime, description, userId)
	if err != nil {
		return -1, err
	}
	fileId, err := filedao.InsertFile(tx, userId, fatherId, uint64(contentId), nowTime)
	if err != nil {
		tx.Rollback()
		return -1, err
	}
	tx.Commit()
	return fileId, err
}
