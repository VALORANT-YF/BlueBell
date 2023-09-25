package routers

import (
	"bluebell/controller"
	"bluebell/middles"

	"github.com/gin-gonic/gin"
)

// 用户相关的controller
var user controller.Users

// UsersGroup 用户的请求路径
func UsersGroup(r *gin.Engine) {
	usersGroup := r.Group("/user")
	{
		usersGroup.POST("/signUp", user.SignUp)
		usersGroup.GET("/login", user.Login)
		usersGroup.GET("/ping", middles.JWTAuthMiddleware(), user.Ping)
	}
}
