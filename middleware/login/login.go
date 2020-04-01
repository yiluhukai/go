package login

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"yiluhuakai/pop/const"
	"yiluhuakai/pop/db/mysql"
	"yiluhuakai/pop/util"

	"fmt"
	"yiluhuakai/logger"
	"yiluhuakai/pop/model"
)

func ValidateLogin() gin.HandlerFunc {
	return func(c *gin.Context) {
		var err error
		defer func() {
			if err != nil {
				c.Set(LoginStatus, false)
			}
		}()

		var loginHeader = &model.LoginHeader{}
		//获取请求头中的信息
		err = c.BindHeader(loginHeader)
		logger.LogDebug("bind header successfully:%#v", loginHeader)
		if err != nil {

			logger.LogError("bind header failed:%v", err)
			util.ResponseError(c, util.ErrCodeParameter)
			c.Abort()
			return
		}

		if len(loginHeader.Code) != 0 && len(loginHeader.ENCRYPTED_DATA) != 0 && len(loginHeader.IV) != 0 {
			var params = model.Login{}
			// 获取third_session
			params.AppId = config.AppId
			params.AppSecret = config.AppSceret
			// 请求微信服务器
			url := fmt.Sprintf(config.AuthCode2Session, params.AppId, params.AppSecret, loginHeader.Code)

			result, err := util.HttpGet(url)
			if err != nil {
				logger.LogError("reuqest tenxun server failed:%v", err)
				util.ResponseError(c, util.ErrCodeServerBusy)
				c.Abort()
				return
			}
			logger.LogDebug("request result = %v", result)

			if result.ErrCode == 0 {
				//设置session
				logger.LogDebug("debug seesion_id =%v ,openid =%v", result.SessionKey, result.OpenId)

				loginHeader.AppId = config.AppId

				loginHeader.SessionKey = result.SessionKey
				userInfoStr, err := loginHeader.AesDecrypt()
				if err != nil {
					logger.LogError("解码失败：%v", err)
					util.ResponseError(c, util.ErrCodeServerBusy)
					c.Abort()
					return
				}

				logger.LogDebug("解密成功")

				// 序列化获取的数据到模型中
				var userInfo = &model.UserInfo{}
				err = json.Unmarshal([]byte(userInfoStr), userInfo)
				if err != nil {
					logger.LogError("failed to json unmarshal:%v ", err)
					util.ResponseError(c, util.ErrCodeServerBusy)
					c.Abort()
					return
				}
				logger.LogDebug("json marshal userInfo successfully")

				user, err := mysql.GetUserInfoByOpenId(userInfo.OpenId)

				if err != nil {
					logger.LogError("get userInfo by openId fail;%v", err)
					util.ResponseError(c, util.ErrCodeServerBusy)
					c.Abort()
					return
				}
				logger.LogDebug("get userInfo successful")

				if user == nil {
					logger.LogDebug("userInfo doesn't exist")
					//用户不存在
					user := &model.User{
						UserId:     util.Gen_uuid(),
						OpenId:     userInfo.OpenId,
						UserInfo:   userInfoStr,
						Sky:        loginHeader.Skey,
						SessionKey: loginHeader.SessionKey,
					}
					err = mysql.SaveUser(user)
				} else {
					logger.LogDebug("userInfo  exist")
					user.UserInfo = userInfoStr
					user.Sky = loginHeader.Skey
					user.SessionKey = loginHeader.SessionKey
					user.OpenId = userInfo.OpenId
					err = mysql.UpdateUser(user)
				}
				if err != nil {
					logger.LogDebug("failed to save or update user", err)
					util.ResponseError(c, util.ErrCodeServerBusy)
					c.Abort()
					return
				}

				logger.LogDebug("save or update user successfully")

				c.Set(LoginStatus, true)
				c.Set(UserInfo, *userInfo)
				c.Set(Sky, loginHeader.Skey)

				c.Next()

			} else {
				logger.LogDebug("fetch openid  and session_key failed,%v", result.ErrMsg)
				util.ResponseError(c, util.ErrCodeServerBusy)
				c.Abort()
				return
			}
			return
		}

		util.ResponseError(c, util.ErrCodeParameter)
		c.Abort()
	}
}
