package userservice

import (
	"github.com/sunlggggg/piconline/main/entity"
	"github.com/sunlggggg/piconline/main/dao/userdao"
	"github.com/sunlggggg/piconline/main/code"
	"log"
	"github.com/sunlggggg/piconline/main/jwtutil"
	"github.com/sunlggggg/piconline/main/config/jwtconfig"
	"time"
)

// 注册一个用户
func Register(user entity.User) (status string, id int64) {
	if userdao.FindByName(user.Name) != nil {
		log.Println(`用户已经存在`)
		return code.UserExisted, -1
	} else {
		user.Id = userdao.Create(user)
		log.Println(user, `创建成功`)
		return code.CreateSuccess, int64(user.Id)
	}
}

func Login(username string, password string) (status string, token string) {
	isRight, err := userdao.CheckPassword(username, password)
	if err != nil {
		log.Println(err)
		return code.LoginErr, ""
	}
	if isRight {
		jwt := jwtutil.JWT{}
		jwt.Header = jwtutil.Header{Alg: "HS256", Typ: "JWT"}
		jwt.PayLoad = jwtutil.PayLoad{ExpireTime: jwtconfig.ValidTime, Name: username, CreateTime: time.Now().Unix()}
		result, _ := jwt.Encode()
		return code.LoginSuccess, result
	} else {
		return code.LoginFail, ""
	}
}
