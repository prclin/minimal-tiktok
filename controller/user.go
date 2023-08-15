package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/request"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/service"
	"net/http"
)

func init() {
	publish := core.ContextRouter.Group("/user")
	publish.POST("/register", RegisterUser)
}

/*
RegisterUser 注册用户

参数 request.UserRegisterRequest

响应 userId token
*/
func RegisterUser(c *gin.Context) {
	urr := request.UserRegisterRequest{}
	if err := c.ShouldBindQuery(&urr); err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.UserEnrollResponse{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "参数错误",
			},
		})
		return
	}
	user := entity.User{
		Username: urr.Username,
		Password: urr.Password,
	}
	//注册逻辑
	registerUser := service.RegisterUser(user)
	c.JSON(http.StatusOK, registerUser)
}
