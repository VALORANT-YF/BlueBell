package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"fmt"
	"net/http"

	"go.uber.org/zap"

	"github.com/gin-gonic/gin"
)

type Users struct{}

const CtxtUserIDKey = "userID"

func (u Users) SignUp(context *gin.Context) {
	//获取参数 参数校验
	var p models.UserRegisterParams
	err := context.ShouldBindJSON(&p)
	fmt.Println(p)
	if err != nil {
		//请求参数错误
		zap.L().Error("SignUp with invalid param", zap.Error(err)) //使用zap记录错误日志
		context.JSON(http.StatusOK, gin.H{
			"msg": err.Error(),
		})
		return
	}
	fmt.Println(p)
	//业务处理
	//传递结构体时,最好传入指针,防止数据过多,影响性能
	err = logic.SignUp(&p)
	if err != nil {
		context.JSON(http.StatusOK, gin.H{
			"msg": "注册失败",
		})
		return
	}
	//返回响应
	context.JSON(http.StatusOK, gin.H{
		"ok": "ok",
	})
}

// Login 用户登录
func (u Users) Login(context *gin.Context) {
	//获取请求参数
	var p models.UserLoginParams
	err := context.ShouldBind(&p)
	if err != nil {
		//请求参数错误
		zap.L().Error("login with invalid param", zap.Error(err)) //使用zap记录错误信息
		context.String(http.StatusOK, "msg:%v", err)
		return
	}
	//判断账号是否存在
	var userId models.UserId
	token, err := logic.Login(&p, &userId)
	if err != nil {
		context.String(http.StatusOK, "登录失败")
		return
	}
	//返回响应
	ResponseSuccess(context, token)
}

// Ping 用户登录状态 通过请求头中是否有 有效的JWT token 即可
// 用户拿到token之后,每次发送请求都会带着请求头
func (u Users) Ping(context *gin.Context) {
	uid, ok := context.Get(CtxtUserIDKey)
	if !ok {
		ResponseError(context, CodeNeedLogin)
		return
	}
	_, ok = uid.(int64) //获取Token中的user_id
	if !ok {
		ResponseError(context, CodeNeedLogin)
		return
	}
	ResponseSuccess(context, "登录成功")
}
