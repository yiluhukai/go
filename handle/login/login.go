package login

import (
	"github.com/gin-gonic/gin"
	"yiluhuakai/logger"
	"yiluhuakai/pop/middleware/login"
	"yiluhuakai/pop/model"
	"yiluhuakai/pop/util"
)

func LoginHandle(c *gin.Context) {
	loginStatus, exist := c.Get(login.LoginStatus)

	if !exist {
		logger.LogDebug("loginStatus doesn't exist")
		util.ResponseError(c, util.ErrCodeServerBusy)
	}

	isLogin, ok := loginStatus.(bool)

	if !ok {
		logger.LogDebug("loginStatus' type is invalid")
		util.ResponseError(c, util.ErrCodeServerBusy)
	}

	if !isLogin {
		logger.LogDebug("user doesn't exist")
		util.ResponseError(c, util.ErrCodeNotLogin)
	}

	// 获取用户的信息并返回
	userInfoIntf, exist := c.Get(login.UserInfo)

	if !exist {
		logger.LogDebug("UserInfo doesn't exist")
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	userInfo, ok := userInfoIntf.(model.UserInfo)

	if !ok {
		logger.LogDebug("userInfo' type is invalid")
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	skyIntf, exist := c.Get(login.Sky)

	if !exist {
		logger.LogDebug("UserInfo doesn't exist")
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	sky, ok := skyIntf.(string)

	if !ok {
		logger.LogDebug("sky's type is invalid")
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	resData := map[string]interface{}{
		"userInfo": userInfo,
		"sky":      sky,
	}

	util.ResponseSuccess(c, resData)

}
