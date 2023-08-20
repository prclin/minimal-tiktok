package util

import (
	"fmt"
	"mime/multipart"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/prclin/minimal-tiktok/global"
)

func getOssClient() *oss.Client {

	for _, env := range os.Environ() {
		fmt.Println(env)
	}

	provider, err := oss.NewEnvironmentVariableCredentialsProvider()
	if err != nil {
		global.Logger.Error(err.Error())
	}

	// 创建OSSClient实例。
	// yourEndpoint填写Bucket对应的Endpoint，以华东1（杭州）为例，填写为https://oss-cn-hangzhou.aliyuncs.com。其它Region请按实际情况填写。
	client, err := oss.New("oss-cn-beijing.aliyuncs.com", "", "", oss.SetCredentialsProvider(&provider))
	if err != nil {
		global.Logger.Error(err.Error())
	}

	return client
}

func UploadToOss(data *multipart.FileHeader) (string, bool) {
	client := getOssClient()

	bucket, err := client.Bucket("vls-tiktok")
	if err != nil {
		global.Logger.Error(err.Error())
		return "", false
	}
	objectName := "video/" + data.Filename

	//文件读取解析
	src, err2 := data.Open()
	if err2 != nil {
		global.Logger.Error(err2.Error())
		return "", false
	}

	err3 := bucket.PutObject(objectName, src)
	if err3 != nil {
		global.Logger.Error(err3.Error())
		return "", false
	}

	return objectName, true

}
