package controller

import (
	"bluebell/logic"
	"strconv"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// 跟社区相关

type CommunityController struct{}

// CommunityHandler 查询到所有跟社区(community_id , community_name) 以列表的形式返回
func (c CommunityController) CommunityHandler(context *gin.Context) {
	data, err := logic.GetCommunityListService()
	if err != nil {
		zap.L().Error("logic.GetCommunity() failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, data)
}

// CommunityDetailHandler 获取社区详情
func (c CommunityController) CommunityDetailHandler(context *gin.Context) {
	//获取社区id
	idStr := context.Param("id") //获取路径参数
	communityId, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		ResponseError(context, CodeInvalidParam)
		return
	}
	//根据id 获取社区详情
	communityDetails, err := logic.GetCommunityDetailService(communityId)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetailService failed", zap.Error(err))
		ResponseError(context, CodeServerBusy)
		return
	}
	ResponseSuccess(context, communityDetails)
}
