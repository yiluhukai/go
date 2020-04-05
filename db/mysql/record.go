package mysql

import (
	"database/sql"
	"yiluhuakai/logger"
	"yiluhuakai/pop/model"
)

func GetLastInsertRecord(openId string) (record *model.Record, err error) {
	record = &model.Record{}
	sqlStr := "select id,openId,mark,create_time,`add` from records where openId =? order by id desc limit ?,?"
	err = db.Get(record, sqlStr, openId, 0, 1)
	if err == sql.ErrNoRows {
		err = nil
		record = nil
		logger.LogDebug("records is empty")
		return
	}
	if err != nil {
		logger.LogError("failed to query in records table:%v", err)
		return
	}
	logger.LogDebug("query records successfully")
	return
}

func InsertRecordIntoRecords(record *model.Record) (err error) {
	sqlStr := "insert into records(`add`,mark,openId) values(?,?,?)"
	_, err = db.Exec(sqlStr, record.Add, record.Mark, record.OpenId)
	if err != nil {
		logger.LogError("insert record into records failed:%v", err)
		return
	}
	logger.LogDebug("record was inserted into table successfully")
	return
}

func DeleteLastRecordById(openId string) (record *model.Record, err error) {
	tx, err := db.Beginx()

	if err != nil {
		logger.LogError("start tx failed:%v", err)
		return
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	sqlStr := "delete from records where openId = ?"
	_, err = tx.Exec(sqlStr, openId)

	if err != nil {
		logger.LogDebug("delete last record failed:%v", err)
		return
	}
	logger.LogDebug("delete successfully")
	record = &model.Record{}
	sqlStr = "select id,openId,mark,create_time,`add` from records where openId =? order by id desc limit ?,?"
	err = tx.Get(record, sqlStr, openId, 0, 1)
	if err == sql.ErrNoRows {
		err = nil
		logger.LogDebug("records is empty")
		err = tx.Commit()
		return
	}
	if err != nil {
		logger.LogError("failed to query in records table:%v", err)
		return
	}
	err = tx.Commit()
	return
}

func GetRecordsByOpenId(openId string, page int64) (record []*model.Record, err error) {
	record = make([]*model.Record, 0)
	sqlStr := "select id,`add`,mark,create_time,openId,note from records where openId =? order by id desc limit ?,?"
	err = db.Select(&record, sqlStr, openId, page*15, 15)

	if err != nil {
		logger.LogError("query recoreds failed:%v", err)
		return
	}
	logger.LogDebug("query records successfully")
	return
}

func UpdateRecordNoteById(record *model.Record) (err error) {
	sqlStr := "update records set note =? where id =?"
	_, err = db.Exec(sqlStr, record.Note, record.Id)

	if err != nil {
		logger.LogError("upate record's note failed:%v", err)
		return
	}
	logger.LogDebug("update record successfully")
	return
}
