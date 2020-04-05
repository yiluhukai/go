package record

import (
	"github.com/gin-gonic/gin"
	"yiluhuakai/logger"
	"yiluhuakai/pop/db/mysql"
	"yiluhuakai/pop/model"
	"yiluhuakai/pop/util"
)

func CreateRecordHandle(c *gin.Context) {
	var record = &model.Record{}

	err := c.BindJSON(record)
	if err != nil {
		logger.LogError("bind params failed:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	logger.LogDebug("bind params successfuly")

	if len(record.OpenId) == 0 || record.Add == 0 {
		logger.LogError("params is invalid:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	//查询最后一条记录

	oldRecord, err := mysql.GetLastInsertRecord(record.OpenId)
	if err != nil {
		logger.LogError("get last record from   records' tables", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	if oldRecord == nil {
		logger.LogDebug("table of records  dosn't have  last record that openid = ? ", record.OpenId)
		record.Mark = record.Add

	} else {

		logger.LogDebug("table of records  have last record that openid = ?", record.OpenId)

		record.Mark = oldRecord.Mark + record.Add
	}
	err = mysql.InsertRecordIntoRecords(record)
	if err != nil {
		logger.LogError("save a new record into table of records failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	logger.LogDebug("insert a new reocod into table of records")

	util.ResponseSuccess(c, nil)
}

func GetRecordsHandle(c *gin.Context) {
	// 获取用户的openid
	openId, exist := c.GetQuery("openId")

	if !exist || len(openId) == 0 {
		logger.LogDebug("openId doesn;t exist")
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	logger.LogDebug("openId = %v", openId)

	page, err := util.GetQueryInt64(c, "page")

	if err != nil {
		logger.LogError("page doen't exist:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	records, err := mysql.GetRecordsByOpenId(openId, page)
	if err != nil {
		logger.LogError("query records failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	logger.LogDebug("query records successfully")

	util.ResponseSuccess(c, records)
}

func GetMarkHandle(c *gin.Context) {
	openId, exist := c.GetQuery("openId")

	if !exist || len(openId) == 0 {
		logger.LogDebug("openId doesn;t exist")
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	logger.LogDebug("openId = %v", openId)

	record, err := mysql.GetLastInsertRecord(openId)

	if err != nil {
		logger.LogError("get last record from   records' tables", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	logger.LogDebug("get mark by openId successfully")

	util.ResponseSuccess(c, *record)
}

func ResetRecordHandle(c *gin.Context) {
	var record = &model.Record{}

	err := c.BindJSON(record)
	if err != nil {
		logger.LogError("bind params failed:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	logger.LogDebug("bind params successfuly")

	if len(record.OpenId) == 0 {
		logger.LogError("params is invalid:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	record.Add = -record.Mark
	record.Mark = 0

	err = mysql.InsertRecordIntoRecords(record)
	if err != nil {
		logger.LogError("save a new record into table of records failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}

	logger.LogDebug("insert a new reocod into table of records")

	util.ResponseSuccess(c, nil)

}

func DeleteLastRecordHandle(c *gin.Context) {
	// 获取用户的openid
	var record = &model.Record{}

	err := c.BindJSON(record)
	if err != nil {
		logger.LogError("bind params failed:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	logger.LogDebug("bind params successfuly")

	if len(record.OpenId) == 0 {
		logger.LogError("params is invalid:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	//delete last records by openId
	record, err = mysql.DeleteLastRecordById(record.OpenId)

	if err != nil {
		logger.LogError("delete last record failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	logger.LogDebug("successfully delete last record bi openId=%v", record.OpenId)

	util.ResponseSuccess(c, *record)
}

func UpdateRecordNoteHandle(c *gin.Context) {
	var record = &model.Record{}
	err := c.BindJSON(&record)

	if err != nil {
		logger.LogError("bind params failed:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}

	logger.LogDebug("bind params successfully")
	if len(record.Note) == 0 {
		logger.LogError("params is invlaid:%v", err)
		util.ResponseError(c, util.ErrCodeParameter)
		return
	}
	err = mysql.UpdateRecordNoteById(record)
	if err != nil {
		logger.LogError("update user failed:%v", err)
		util.ResponseError(c, util.ErrCodeServerBusy)
		return
	}
	util.ResponseSuccess(c, nil)
}
