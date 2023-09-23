package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	"bluebell/pkg/jwtToken"
	"bluebell/pkg/snowflake"
	"errors"
	"fmt"
)

// SignUp 用户注册
func SignUp(p *models.UserRegisterParams) (err error) {
	//判断用户是否存在
	isExists, err := mysql.QueryByUsername(p.Username)
	if err != nil {
		//数据库查询出错
		return err
	}
	if isExists {
		return errors.New("用户已经存在")
	}
	//生成user_id
	userID := snowflake.GenID()
	//构造一个User实例
	u := models.User{
		UserId:   userID,
		Username: p.Username,
		Password: p.Password,
	}
	//将用户注册的信息保存到数据库
	err = mysql.InsertUser(&u)
	return
}

// Login 用户登录
func Login(p *models.UserLoginParams, u *models.UserId) (token string, err error) {
	//查询用户user_id
	userIdExists, err := mysql.QueryUserId(p, u)
	if err != nil {
		return "", err
	}
	//如果查询到id 说明登录成功
	if userIdExists {
		//生成JWT的Token
		aToken, _, _ := jwtToken.GenToken(u.UserId)
		fmt.Println("@@@", aToken)
		return token, nil
	}
	return "", err
}
