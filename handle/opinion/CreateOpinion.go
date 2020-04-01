package opinion

import (
	"yiluhuakai/pop/db/mysql"

	"github.com/gin-gonic/gin"
	"yiluhuakai/logger"
	"yiluhuakai/pop/model"
	"yiluhuakai/pop/util"
)

func CreateOponionHandle(c *gin.Context) {
	var opinion = &model.Opinion{}
	err := c.BindJSON(opinion)

	if err != nil {
		logger.LogError("bind params failed;%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	logger.LogDebug("opinion = %v", opinion)

	// 保存到数据库中
	err = mysql.CreateOpinion(opinion)
	if err != nil {
		logger.LogError("insert into oponion database failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	logger.LogDebug("insert into opinion database successfully;%v")

	util.ResponseSuccess(c, nil)
}
