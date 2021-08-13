package dao

import (
	"TTMS/model"
	"TTMS/utils"
)

func AddSession(sess *model.Session) error {
	//写sql语句
	sqlStr := "insert into sessions values(?,?,?)"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, sess.SessionID, sess.UserName, sess.UserID)
	if err != nil {
		return err
	}
	return nil
}
func DeleteSession(sessID string) error {
	//写sql语句
	sqlStr := "delete from sessions where session_id = ?"
	//执行sql
	_, err := utils.Db.Exec(sqlStr, sessID)
	if err != nil {
		return err
	}
	return nil
}