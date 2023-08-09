package core

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/prclin/minimal-tiktok/global"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/model/token"
	"net/http"
	"strconv"
	"strings"
	"time"
)

var Router *gin.Engine
var ContextRouter *gin.RouterGroup

/*
initGin 初始化gin
*/
func initGin() {
	//创建gin engine
	engine := gin.New()
	//注册全局中间件，使用自定义的日志中间件，使用gin默认的recover中间件
	engine.Use(parse(), ginLogger(), gin.Recovery())
	//404处理
	engine.NoRoute(func(context *gin.Context) {
		context.JSON(404, response.Response{StatusCode: 404, StatusMsg: "Not Found!"})
	})
	Router = engine
	Router = engine
	ContextRouter = engine.Group(global.Configuration.Server.ContextPath)
}

/*
ginLogger 是自定义的全局日志中间件，用于替代gin的默认日志中间件
*/
func ginLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		cost := time.Since(start) // 本次请求的总共消耗时间
		// 写入日志
		global.Logger.Infow(
			strconv.Itoa(c.Writer.Status()),
			"path", c.Request.URL.String(),
			"method", c.Request.Method,
			"clientIp", c.ClientIP(),
			"errors", c.Errors.ByType(gin.ErrorTypePrivate).String(),
			"cost", cost,
		)
	}
}

//自定义中间件，用来进行token的校验

func parse() gin.HandlerFunc {
	return func(c *gin.Context) {
		if strings.Contains(c.Request.RequestURI, "/douyin/user/register") || strings.Contains(c.Request.RequestURI, "/douyin/user/login") {
			return
		}

		//测试解析功能
		//校验

		if tt, ok := c.Request.Header["Token"]; ok {

			tn, err := jwt.ParseWithClaims(tt[0], &token.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte("1234"), nil
			})
			if err != nil {
				var res = response.Douyin_user_register_response{
					StatusCode: -1,
					StatusMsg:  "token以过期,请重新登录！",
					UserId:     -1,
					Token:      "",
				}
				c.AbortWithStatusJSON(http.StatusOK, res)
				return
			}

			//判断校验结果

			if _, ok := tn.Claims.(*token.MyCustomClaims); ok && tn.Valid {
				c.Next()

			} else {
				var res = response.Douyin_user_register_response{
					StatusCode: -1,
					StatusMsg:  "token以过期,请重新登录！",
					UserId:     -1,
					Token:      "",
				}
				c.JSON(http.StatusOK, res)
			}

		}

	}

}
