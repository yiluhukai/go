package upload

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"yiluhuakai/logger"
	"yiluhuakai/pop/util"
)

func UploadFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		logger.LogError("upload file failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	logger.LogDebug("file = %#v", file.Filename)
	//保存文件到本地
	//err = c.SaveUploadedFile(file, fileName)

	//上传到腾讯云

	File, err := file.Open()

	if err != nil {
		logger.LogError("open file failed;%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	presignedURL, err := util.GetPresignedURL(file.Filename)
	if err != nil {
		logger.LogError("get PresignedURL failed;%v", err)
		panic(err)
	}
	req, err := http.NewRequest(http.MethodPut, presignedURL.String(), File)
	if err != nil {
		logger.LogError("create a new request failed;%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	//req.Header.Set("Content-Type", "image/webp")

	_, err = http.DefaultClient.Do(req)

	if err != nil {
		logger.LogError("upload file to tenxun failed;%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	logger.LogDebug("upload file successfully")
	// logger.LogDebug("header =%#v", res.Header)

	// body, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	logger.LogError("parse response body failed:%v", err)
	// 	util.ResponseError(c, util.ErrCodeServerBusy)
	// 	return
	// }
	//logger.LogDebug("reponse content:%v", body)
	// logger.LogDebug("EscapedPath =%v", presignedURL.EscapedPath())
	// logger.LogDebug("Hostname=%v", presignedURL.Hostname())
	// logger.LogDebug("Path =%v", presignedURL.Path)
	// logger.LogDebug("RequestURI=%v", presignedURL.RequestURI())

	// logger.LogDebug("Scheme=%v", presignedURL.Scheme)
	url := presignedURL.String()
	index := strings.Index(url, "?")
	data := []byte(url)
	util.ResponseSuccess(c, string(data[0:index]))

}
