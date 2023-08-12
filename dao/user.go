package dao

import (
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
)

func GetUserById(id int) (u entity.User) {
	var user entity.User
	global.Datasource.Where("id = ?", id).First(&user)

	return user
}
func CreateUser(user *entity.User) {

	global.Datasource.Create(user)

}
func CreateUserDto(userDto *entity.UserDto) {

	global.Datasource.Create(userDto)

}
