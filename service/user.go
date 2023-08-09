package service

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/prclin/minimal-tiktok/dao"
	"github.com/prclin/minimal-tiktok/model/response"
	"github.com/prclin/minimal-tiktok/model/token"
	"net/http"
	"time"
)

func RegisterHandler(c *gin.Context) {

	//获取用户名和密码
	userName := c.Query("username")
	password := c.Query("password")
	//初始化redis用来进行token的存储
	//client := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "",
	//	DB:       0,
	//})
	//ctx := context.Background()
	//查询该用户是否已经注册过
	user := dao.GetUserByName(userName)
	if user.Username != "" {
		//若已经注册则跳转到登录页面
		var res response.Douyin_user_register_response
		res.StatusCode = 200
		res.StatusMsg = "已经注册过了"
		res.UserId = -1

		c.JSON(http.StatusOK, res)
	} else {
		//若没有注册进行创建用户，并登录添加token
		user.Username = userName
		user.Password = password
		user.CreateTime = time.Now()
		user.Extra = "{}"
		dao.CreateUser(&user)
		//生成token，使用jwt框架
		claims := token.MyCustomClaims{
			userName,
			password,
			jwt.RegisteredClaims{
				// A usual scenario is to set the expiration time relative to the current time
				ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), //过期时间为一天
				IssuedAt:  jwt.NewNumericDate(time.Now()),
				NotBefore: jwt.NewNumericDate(time.Now()),
				Issuer:    "admin",
				//Subject:   "somebody",
				//ID:        "1",
				//Audience:  []string{"somebody_else"},
			},
		}
		t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		//ss为生成的token
		ss, _ := t.SignedString([]byte("1234"))
		//使用redis存储token
		//client.Set(ctx, strconv.FormatInt(user.ID, 10), ss, 24*time.Hour)

		//////测试解析功能
		//////校验
		//s2, err := client.Get(ctx, strconv.FormatInt(user.ID, 10)).Result()
		//if err != nil {
		//	var res = response.Douyin_user_register_response{
		//		StatusCode: -1,
		//		StatusMsg:  "token以过期,请重新登录！",
		//		UserId:     -1,
		//		Token:      "",
		//	}
		//
		//	c.JSON(http.StatusOK, res)
		//}
		//tn, _ := jwt.ParseWithClaims(s2, &token.MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		//	return []byte("1234"), nil
		//})
		////判断校验结果
		//
		//if claims, ok := tn.Claims.(*token.MyCustomClaims); ok && tn.Valid {
		//	fmt.Printf("%v************%v", claims.UserName, claims.Password)
		//} else {
		//	fmt.Println(ok, "++++++++++++++++++++++++++++++++++++++buok")
		//}

		//设置返回信息
		var res response.Douyin_user_register_response
		res.StatusCode = 0
		res.StatusMsg = "成功"
		res.UserId = user.ID
		res.Token = ss

		c.JSON(http.StatusOK, res)

	}

}
