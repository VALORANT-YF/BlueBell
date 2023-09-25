package controller

import (
	"bluebell/logic"
	"bluebell/models"
	"errors"
	"strconv"

	"github.com/go-playground/validator/v10"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

type Post struct{}

var UserNoLogin = errors.New("用户未登录")

// CreatePostHandler 创建贴子
func (p Post) CreatePostHandler(context *gin.Context) {
	post := new(models.Post)
	//1.获取参数
	err := context.ShouldBindJSON(post)
	if err != nil {
		zap.L().Error("create post with  invalid param")
		ResponseError(context, CodeInvalidParam)
		return
	}
	//从Token中获取当前登录用户的id
	uid, ok := context.Get(CtxtUserIDKey)
	if !ok {
		err = UserNoLogin
		return
	}
	userId, ok := uid.(int64)
	if !ok {
		err = UserNoLogin
		return
	}
	post.AuthorID = userId
	//2.创建贴子
	if err := logic.CreatePost(post); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, nil)
}

// GetPostDetailHandler 得到帖子详情
func (p Post) GetPostDetailHandler(context *gin.Context) {
	//获取参数 帖子id
	postIdStr := context.Param("id")
	//将字符串转化为int64
	postId, err := strconv.ParseInt(postIdStr, 10, 64)
	if err != nil {
		zap.L().Error("GetPostDetailHandler param is Error", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}
	//根据帖子id 获取帖子详情的处理函数
	data, err := logic.GetPostDetailById(postId)
	if err != nil {
		zap.L().Error("logic.GetPostDetailHandler", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}
	ResponseSuccess(context, data)
}

// GetPostListDetailHandler 帖子分页
func (p Post) GetPostListDetailHandler(context *gin.Context) {
	//获取分页请求参数
	postListParams := new(models.PostList)
	err := context.ShouldBind(&postListParams)
	if err != nil {
		zap.L().Error("postListParams is Error", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}
	//获取数据
	data, err := logic.GetPostList(postListParams)
	if err != nil {
		zap.L().Error("logic.GetPostListDetailHandler", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, data)
}

// PostVoteHandler 投票
func (p Post) PostVoteHandler(context *gin.Context) {
	//参数校验
	vote := new(models.VoteData)
	if err := context.ShouldBind(vote); err != nil {
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs) //类型断言
		if !ok {
			ResponseError(context, CodeInvalidParam)
			return
		}
		ResponseErrorWithMsg(context, CodeInvalidParam, "投票失败")
		return
	}
	//获取当前投票用户的id
	//从Token中获取当前登录用户的id
	uid, ok := context.Get(CtxtUserIDKey)
	if !ok {
		zap.L().Error("context.Get(CtxtUserIDKey) is error")
		return
	}
	userId, ok := uid.(int64)
	if !ok {
		zap.L().Error("uid.(int64) is error")
		return
	}
	err := logic.PostVote(userId, vote)
	if err != nil {
		zap.L().Error("logic.VoteForPost() failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, nil)
}

// GetPostListDetailHandler2 根据时间或分数排序帖子
func (p Post) GetPostListDetailHandler2(context *gin.Context) {
	//获取分页请求参数
	postListParams := &models.ParamPostList{
		Page:  1,
		Size:  10,
		Order: "time",
	}
	err := context.ShouldBindQuery(postListParams)
	if err != nil {
		zap.L().Error("postListParams is Error", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}
	//获取数据
	data, err := logic.GetPostList2(postListParams)
	if err != nil {
		zap.L().Error("logic.GetPostListDetailHandler", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, data)
}

// GetCommunityPostListHandler 根据社区查询帖子列表
func (p Post) GetCommunityPostListHandler(context *gin.Context) {
	//获取分页请求参数
	postListParams := &models.ParamCommunityPostList{
		ParamPostList: &models.ParamPostList{
			Page:  1,
			Size:  10,
			Order: "time",
		},
	}
	err := context.ShouldBindQuery(postListParams)
	if err != nil {
		zap.L().Error("ParamCommunityPostList is Error", zap.Error(err))
		ResponseError(context, CodeInvalidParam)
		return
	}
	//获取数据
	data, err := logic.GetCommunityPostList2(postListParams)
	if err != nil {
		zap.L().Error("logic.GetPostListDetailHandler", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, data)
}
