package service

// import (
// 	"log"
// 	"math/rand"
// 	"mime/multipart"
// 	"strings"
// 	"time"

// 	"github.com/minio/minio-go/v6"
// 	"github.com/prclin/minimal-tiktok/global"
// )

// // 生成随机字符串的字符集
// const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// func getMinIoClient() *minio.Client {
// 	endpoint := "192.168.200.129:9000"
// 	accessKeyID := "admin"
// 	secretAccessKey := "admin321"
// 	useSSL := false
// 	client, err := minio.New(endpoint, accessKeyID, secretAccessKey, useSSL)
// 	if err != nil {
// 		log.Fatalln("创建 MinIO 客户端失败", err)
// 		return nil
// 	}

// 	return client
// }

// func uploadVideo(client *minio.Client, data *multipart.FileHeader) (string, bool) {
// 	bucketName := "tiktok"
// 	// 目标 bucket 的位置
// 	location := ""
// 	//创建桶
// 	exists, err := client.BucketExists(bucketName)
// 	if err != nil {
// 		global.Logger.Error(err.Error())
// 	}
// 	if !exists {
// 		err = client.MakeBucket(bucketName, location)
// 		if err != nil {

// 			global.Logger.Error(err.Error())
// 		}
// 	}

// 	// 指定上传视频名
// 	objectName := data.Filename
// 	// 指定上传文件类型
// 	contentType := "video/mp4"
// 	//文件读取解析
// 	src, err2 := data.Open()
// 	if err2 != nil {
// 		global.Logger.Error(err2.Error())
// 		return "", false
// 	}
// 	defer src.Close()

// 	//小问题，可能上传不了128MB的文件
// 	//文件上传
// 	_, err = client.PutObject(bucketName, objectName, src, -1, minio.PutObjectOptions{
// 		ContentType: contentType,
// 	})
// 	if err != nil {
// 		global.Logger.Error(err.Error())
// 		return "", false
// 	}

// 	path := "/minio/" + bucketName + "/" + objectName

// 	return path, true
// }

// /*
// *

// 	生成随机字符串

// *
// */
// func randomString(n int) string {

// 	rand.Seed(time.Now().UnixNano())

// 	sb := strings.Builder{}
// 	sb.Grow(n)
// 	for i := 0; i < n; i++ {
// 		sb.WriteByte(charset[rand.Intn(len(charset))])
// 	}
// 	return sb.String()

// }
