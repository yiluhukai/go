package mysql

import (
	"yiluhuakai/logger"
	"yiluhuakai/pop/model"
)

func CreateOpinion(opon *model.Opinion) (err error) {

	sqlStr := "insert into opinion(openId,opinion,wechat,src) values(?,?,?,?)"
	_, err = db.Exec(sqlStr, opon.OpenId, opon.Opinion, opon.WeChat, opon.Src)

	if err != nil {
		logger.LogError("create a oponion failed;%v ", err)
		return
	}
	logger.LogDebug("crate oponion successfully")
	return
}
