package main

import (
	"github.com/gin-gonic/gin"
	"yiluhuakai/logger"
	"yiluhuakai/pop/db/mysql"
	"yiluhuakai/pop/handle/login"
	"yiluhuakai/pop/handle/opinion"
	"yiluhuakai/pop/handle/record"
	"yiluhuakai/pop/handle/upload"
	"yiluhuakai/pop/middleware/checkSession"
	middlewareLogin "yiluhuakai/pop/middleware/login"
	"yiluhuakai/pop/util"
)

func init() {
	// 初始化日志
	err := logger.InitLogger("console", "", "", logger.LogLevelDebug, "")

	if err != nil {
		logger.LogError("connect mysql database failed:%v", err)
		panic(err)
	}

	//初始化数据库路

	mysql.InitDb()

	err = util.NewClient()

	if err != nil {
		logger.LogError("create cos client failed;%v", err)
		panic(err)
	}

}

func main() {
	router := gin.Default()

	router.Static("/upload", "./upload")

	authGroup := router.Group("/weapp/", middlewareLogin.ValidateLogin())
	authGroup.GET("/login", login.LoginHandle)

	apiGroup := router.Group("/api/", checkSession.CheckSessionMiddleWare())
	apiGroup.POST("/upload", upload.UploadFile)
	apiGroup.POST("/create_opinion", opinion.CreateOponionHandle)
	apiGroup.POST("/createrecord", record.CreateRecordHandle)
	apiGroup.POST("/resetmark", record.ResetRecordHandle)
	apiGroup.POST("/deleterecord", record.DeleteLastRecordHandle)
	apiGroup.POST("/updatenote", record.UpdateRecordNoteHandle)
	apiGroup.GET("/getmark", record.GetMarkHandle)
	apiGroup.GET("/getrecords", record.GetRecordsHandle)
	router.Run() // 监听并在 0.0.0.0:8080 上启动服务
}
