package service

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/util"
)

func Message(c *gin.Context) {

	//检查token
	token := c.Query("token")
	if !util.CheckToken(token) {
		c.JSON(200, gin.H{
			"status_code": 1,
			"status_msg":  "token已失效",
			"user":        nil,
		})
		c.Abort()
		return
	}

	user_id_str := c.Query("user_id")
	//字符串转int
	user_id, _ := strconv.Atoi(user_id_str)
	//根据id查询用户信息
	user_msg := dao.GetUserById(user_id)
	//将user拷贝到userDto
	var userDto entity.UserDto
	util.Copy(user_msg).To(userDto)
	userDto.IsFollow = false

	//返回
	var res response.MessageResponse
	res.StatusCode = 0
	res.StatusMsg = "查询成功"
	res.User = userDto

	c.JSON(http.StatusOK, res)
}
