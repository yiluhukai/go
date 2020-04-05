package util

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
	"yiluhuakai/logger"
)

// 将参数中的字符串转成整形
func GetQueryInt64(c *gin.Context, key string) (value int64, err error) {
	tempValue, ok := c.GetQuery(key)
	if !ok {
		logger.LogDebug("doesn't found value of key")
		err = fmt.Errorf("invalid params,not found key")
		return
	}
	value, err = strconv.ParseInt(tempValue, 10, 64)

	if err != nil {
		logger.LogError("invalid params, strconv.ParseInt failed, err:%v, str:%v", err, tempValue)
		return
	}
	return
}
