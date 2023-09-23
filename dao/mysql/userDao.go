package mysql

import (
	"bluebell/models"
	"bluebell/pkg/md5Password"
)

//把每一步数据库操作封装成函数
//等待service层根据业务需求调用

// QueryByUsername 根据用户名,查询用户信息
func QueryByUsername(username string) (bool, error) {
	sqlStr := `select count(user_id) from user where username = ?`
	var count int
	err := db.Get(&count, sqlStr, username)
	if err != nil {
		return false, err
	}
	//count > 0 存在 返回true
	return count > 0, nil
}

// InsertUser 用户注册
func InsertUser(user *models.User) (err error) {
	password := md5Password.EncryptPassword(user.Password)
	//执行sql语句 插入数据
	sqlStr := `insert into user(user_id , username , password) values (?,?,?)`
	_, err = db.Exec(sqlStr, user.UserId, user.Username, password)
	return
}

// QueryUserId 根据用户输入的信息 查询用户id 判断用户是否存在
func QueryUserId(p *models.UserLoginParams, u *models.UserId) (bool, error) {
	sqlStr := `select  id , user_id from user where username = ? and password = ?`
	//执行sql
	password := md5Password.EncryptPassword(p.Password)

	err := db.Get(u, sqlStr, p.Username, password)

	if err != nil {
		return false, err
	}
	return u.Id > 0, nil
}
