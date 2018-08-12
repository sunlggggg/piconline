package userservice

import (
	"github.com/sunlggggg/piconline/main/entity"
	"github.com/sunlggggg/piconline/main/dao/userdao"
	"github.com/sunlggggg/piconline/main/code"
	"log"
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
