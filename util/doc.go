/*
Package util 存放所有的工具函数
*/
package util

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/prclin/minimal-tiktok/model/token"
	"time"
)

func CreateToken(userName string, password string) string {
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
	return ss
}
