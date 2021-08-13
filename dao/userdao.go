package dao

import (
	"TTMS/model"
	"TTMS/utils"
	"net/http"
)


func CheckUserNameAndPassword(username string, password string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id,username,password,email from users where username = ? and password = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username, password)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email)
	return user, nil
}
func IsLogin(r *http.Request) (bool, *model.Session) {
	//根据Cookie的name获取Cookie
	cookie, _ := r.Cookie("user")
	if cookie != nil {
		//获取Cookie的value
		cookieValue := cookie.Value
		//根据cookieValue去数据库中查询与之对应的Session
		session, _ := GetSession(cookieValue)
		if session.UserID> 0 {
			//已经登录
			return true, session
		}
	}
	//没有登录
	return false, nil
}
func GetSession(sessID string) (*model.Session, error) {
	//写sql语句
	sqlStr := "select session_id,username,user_id from sessions where session_id = ?"
	//预编译
	inStmt, err := utils.Db.Prepare(sqlStr)
	if err != nil {
		return nil, err
	}
	//执行
	row := inStmt.QueryRow(sessID)
	//创建Session
	sess := &model.Session{}
	//扫描数据库中的字段值为Session的字段赋值
	row.Scan(&sess.SessionID, &sess.UserName, &sess.UserID)
	return sess, nil
}
func CheckUserName(username string) (*model.User, error) {
	//写sql语句
	sqlStr := "select id,username,password,email,number from users where username = ?"
	//执行
	row := utils.Db.QueryRow(sqlStr, username)
	user := &model.User{}
	row.Scan(&user.ID, &user.Username, &user.Password, &user.Email,&user.Number)
	return user, nil
}
func SaveUser(username string, password string, email string,number string) error {
	//写sql语句
	sqlStr := "insert into users(username,password,email,number ) values(?,?,?,?)"
	//执行
	_, err := utils.Db.Exec(sqlStr, username, password, email,number)
	if err != nil {
		return err
	}
	return nil
}