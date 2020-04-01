package checkSession

import (
	"github.com/gin-gonic/gin"
	"yiluhuakai/logger"
	"yiluhuakai/pop/db/mysql"
	"yiluhuakai/pop/util"
)

func CheckSessionMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		sky := c.GetHeader(WxAppSky)
		if len(sky) == 0 {
			logger.LogDebug("params is invalid!")
			util.ResponseError(c, util.ErrCodeParameter)
			c.Abort()
			return
		}
		isExist, err := mysql.CheckSkyExist(sky)

		if err != nil {
			logger.LogError("query userInfo via sky failed:%v", err)
			util.ResponseError(c, util.ErrCodeServerBusy)
			c.Abort()
			return
		}
		if !isExist {
			logger.LogDebug("user doesn't exist:%v")
			util.ResponseError(c, util.ErrCodeNotLogin)
			c.Abort()
			return
		}
		logger.LogDebug("user has logined")
		c.Next()
	}
}
