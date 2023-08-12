/*
主要用于检查token是否正确，是否过期
*/
package util

import (
	"time"
)

func CheckToken(token string) bool {
	if token == "" {
		return false
	}
	claims, err := ParseToken(token)
	if err != nil {
		return false
	}
	if claims.ExpiresAt.Time.Unix() < time.Now().Unix() {
		return false
	}
	return true
}
