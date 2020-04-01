package mysql

import (
	"database/sql"
	"yiluhuakai/logger"
	"yiluhuakai/pop/model"
)

func GetUserInfoByOpenId(openId string) (user *model.User, err error) {
	user = &model.User{}
	sqlStr := "select id,user_id,open_id,userInfo,session_key,sky  from userInfo where open_id = ?"
	err = db.Get(user, sqlStr, openId)
	if err == sql.ErrNoRows {
		logger.LogDebug("user doen't exist")
		err = nil
		user = nil
		return
	}
	if err != nil {
		logger.LogError("query userInfo failed;%v", err)
		return
	}
	logger.LogDebug("fetch UserInfo successfully")
	return
}

func SaveUser(user *model.User) (err error) {
	sqlStr := "insert into userInfo(user_id,sky,session_key,userInfo,open_id) values(?,?,?,?,?)"
	logger.LogDebug("userInfo:%v,open_id;%v", len(user.UserInfo), user.OpenId)

	_, err = db.Exec(sqlStr, user.UserId, user.Sky, user.SessionKey, user.UserInfo, user.OpenId)

	if err != nil {
		logger.LogError("save user failed:%v", err)
		return
	}

	logger.LogDebug("save user success")
	return
}

func UpdateUser(user *model.User) (err error) {
	sqlStr := "update userInfo set sky= ?,session_key=?,UserInfo=? where open_id =?"

	_, err = db.Exec(sqlStr, user.Sky, user.SessionKey, user.UserInfo, user.OpenId)

	if err != nil {
		logger.LogError("update user failed:%v", err)
		return
	}

	logger.LogDebug("upadte user success")
	return
}

func CheckSkyExist(sky string) (isExist bool, err error) {
	var user = &model.User{}
	sqlStr := "select id,user_id,open_id,userInfo,session_key,sky from userInfo where sky =?"
	err = db.Get(user, sqlStr, sky)
	if err == sql.ErrNoRows {
		isExist = false
		err = nil
		return
	}
	if err != nil {
		logger.LogError("query userInfo by sky fail :%v", err)
		isExist = false
		return
	}

	logger.LogDebug("query userInfo successfully")
	isExist = true
	return
}
