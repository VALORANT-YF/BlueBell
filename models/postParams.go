package models

//post路由的参数

// PostList 定义post 分页参数
type PostList struct {
	Offset int32 `json:"offset" form:"offset"`
	Limit  int32 `json:"limit" form:"limit"`
}

// VoteData 投票
type VoteData struct {
	//UserID 从请求中获取当前的用户
	PostId    int64 `json:"post_id,string" binding:"required"`       // 帖子 id
	Direction int8  `json:"direction,string" binding:"oneof=1 0 -1"` // 赞成票/反对票/取消投票
}

// ParamPostList 获取帖子列表按照时间或分数排序的参数
type ParamPostList struct {
	CommunityId int64  `json:"community_id" form:"community_id"`
	Page        int64  `json:"page"`
	Size        int64  `json:"size"`
	Order       string `json:"order"`
}

// ParamCommunityPostList 获取帖子列表按照时间或分数排序的参数
type ParamCommunityPostList struct {
	*ParamPostList
}
