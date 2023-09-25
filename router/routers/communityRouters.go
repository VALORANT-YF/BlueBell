package routers

import (
	"bluebell/controller"

	"github.com/gin-gonic/gin"
)

var community controller.CommunityController

func CommunityGroup(r *gin.Engine) {
	communityGroup := r.Group("/api")
	{
		communityGroup.GET("/community", community.CommunityHandler)
		//根据id 查询社区详情
		communityGroup.GET("/communityDetail/:id", community.CommunityDetailHandler)
	}
}
