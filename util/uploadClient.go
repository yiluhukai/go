package util

import (
	"context"
	"github.com/tencentyun/cos-go-sdk-v5"
	"net/http"
	"net/url"
	"time"
	"yiluhuakai/logger"
)

const (
	AppId      = "1253773563"
	SecretId   = "AKIDwhJkxwfcDajeLqrH7h4C6kAMPDl3dybM"
	SecretKey  = "iJekbFhtpX9XhbHjsmQyJPM5asxnZ9h2"
	UploadPath = "https://pop-1253773563.cos.ap-shanghai.myqcloud.com"
)

var (
	Client *cos.Client
)

// 商城cos go客户端

func NewClient() (err error) {
	// 将 examplebucket-1250000000 和 COS_REGION修改为真实的信息
	uploadUrl, err := url.Parse(UploadPath)
	if err != nil {
		logger.LogError("upload url is invalid;%v", err)
		return
	}
	bucket := &cos.BaseURL{BucketURL: uploadUrl}
	// 1.永久密钥
	Client = cos.NewClient(bucket, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "COS_SECRETID",
			SecretKey: "COS_SECRETKEY",
		},
	})
	return
}

func GetPresignedURL(name string) (presignedURL *url.URL, err error) {

	ctx := context.Background()
	// 获取预签名URL
	presignedURL, err = Client.Object.GetPresignedURL(ctx, http.MethodPut, name, SecretId, SecretKey, time.Hour, nil)
	if err != nil {
		logger.LogError("get presinged Url failed:%v", err)
		return
	}
	return
}
