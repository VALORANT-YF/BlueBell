package routers

import (
	"bluebell/controller"
	"bluebell/middles"

	"github.com/gin-gonic/gin"
)

// 帖子的controller
var post controller.Post

func PostGroup(r *gin.Engine) {
	postGroups := r.Group("/post")
	{
		//创建帖子
		postGroups.POST("/create", middles.JWTAuthMiddleware(), post.CreatePostHandler)
		//帖子详情
		postGroups.GET("/postDetail/:id", post.GetPostDetailHandler)
		//帖子列表的分页
		postGroups.GET("/postList", post.GetPostListDetailHandler)
		// 根据时间或分数获取帖子列表
		postGroups.GET("/postList2", post.GetPostListDetailHandler2)
		//投票
		postGroups.POST("/vote", middles.JWTAuthMiddleware(), post.PostVoteHandler)
		postGroups.GET("/communityList", post.GetCommunityPostListHandler)
	}
}
