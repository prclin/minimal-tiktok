package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/prclin/minimal-tiktok/core"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/entity"
	"github.com/prclin/minimal-tiktok/model/request"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/service"
	"github.com/prclin/minimal-tiktok/util"
	"net/http"
	"strconv"
)

func init() {
	user := core.ContextRouter.Group("/user")
	user.POST("/register", RegisterUser)
	user.POST("/login", UserLogin)
	user.GET("", GetUserInfo)
}

/*
RegisterUser 注册用户

参数 request.UserEnrollRequest

响应 userId token
*/
func RegisterUser(c *gin.Context) {
	var uer request.UserEnrollRequest
	if err := c.ShouldBindQuery(&uer); err != nil {
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
		Username: uer.Username,
		Password: uer.Password,
	}
	//注册逻辑
	registerUser := service.RegisterUser(user)
	c.JSON(http.StatusOK, registerUser)
}

/*
UserLogin 用户登录
参数 request.UserEnrollRequest
响应 userId token
*/
func UserLogin(c *gin.Context) {
	//参数校验
	var uer request.UserEnrollRequest
	if err := c.ShouldBindQuery(&uer); err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.UserEnrollResponse{
			Response: response.Response{
				StatusCode: 1,
				StatusMsg:  "参数错误",
			},
		})
		return
	}

	//登录
	user := entity.User{
		Username: uer.Username,
		Password: uer.Password,
	}

	userLogin := service.UserLogin(user)
	c.JSON(http.StatusOK, userLogin)
}

/*
GetUserInfo 获取用户信息

参数 userId token

返回 response.UserInfoResponse
*/
func GetUserInfo(c *gin.Context) {
	//参数校验
	var query struct {
		UserId string `form:"user_id" binding:"required"`
		Token  string `form:"token" binding:"required"`
	}
	err := c.ShouldBindQuery(&query)
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.UserInfoResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "参数错误"},
			User:     nil,
		})
		return
	}
	claims, err := util.ParseToken(query.Token) //token解析
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.UserInfoResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "参数错误"},
			User:     nil,
		})
		return
	}
	erId := claims.Id                                    //获取查询者id
	eeId, err := strconv.ParseUint(query.UserId, 10, 64) //获取被查询者id
	if err != nil {
		global.Logger.Debug(err.Error())
		c.JSON(http.StatusOK, response.UserInfoResponse{
			Response: response.Response{StatusCode: 1, StatusMsg: "参数错误"},
			User:     nil,
		})
		return
	}

	//获取用户信息
	userInfo := service.GetUserInfo(erId, eeId)
	c.JSON(http.StatusOK, userInfo)
}
