package util

import (
	"crypto/md5"
	"crypto/rsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/pem"
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/prclin/minimal-tiktok/global"
	"time"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	Id uint64
}

// MD5 加密字符串
func MD5(raw string) string {
	hash := md5.New()
	sum := hash.Sum([]byte(raw))
	return hex.EncodeToString(sum[:])
}

// GenerateToken 生成token
func GenerateToken(id uint64) (string, error) {
	claims := TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			//过期时间
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
		Id: id,
	}
	//加密
	key, err := parsePrivateKey([]byte(global.Configuration.Jwt.RSA.PrivateKey))
	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)
	return token, err
}

// parsePrivateKey 解析PKCS1格式私钥
func parsePrivateKey(buf []byte) (*rsa.PrivateKey, error) {
	p := &pem.Block{}
	p, buf = pem.Decode(buf)
	if p == nil {
		return nil, errors.New("parse private key error")
	}
	return x509.ParsePKCS1PrivateKey(p.Bytes)
}
