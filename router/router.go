package router

import (
	"bluebell/logger"
	"bluebell/router/routers"

	"github.com/gin-gonic/gin"
)

func Setup() *gin.Engine {
	// 创建一个路由引擎
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 路由分组 加载路由文件
	routers.UsersGroup(r)
	routers.CommunityGroup(r)
	routers.PostGroup(r)

	return r
}
