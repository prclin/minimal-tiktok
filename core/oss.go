package core

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/prclin/minimal-tiktok/global"
	"os"
)

func initOSS() {
	ossConfig := global.Configuration.OSS

	//创建Client实例
	client, err := oss.New(ossConfig.EndPoint, ossConfig.AccessKeyId, ossConfig.AccessKeySecret)
	if err != nil {
		fmt.Printf("创建oss客户端时发生错误：%v\n", err)
		os.Exit(-1)
	}
	global.OSSClient = client
}
