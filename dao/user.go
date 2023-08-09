package dao

import (
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
)

func GetUserByName(userName string) (u entity.User) {
	var user entity.User
	global.Datasource.Where("username= ?", userName).First(&user)

	return user
}
func CreateUser(user *entity.User) {

	global.Datasource.Create(user)

}
