package models

import "time"

type Post struct {
	PostId      int64     `json:"post_id,string" db:"post_id"` //json:"json:"post_id,string" 加上,string 防止数据过大,前端无法处理
	AuthorID    int64     `json:"author_id" db:"author_id"`
	CommunityId int64     `json:"community_id" db:"community_id" binding:"required"`
	Status      int32     `json:"status" db:"status"`
	Title       string    `json:"title" db:"title" binding:"required"`
	Content     string    `json:"content" db:"content" binding:"required"`
	CreateTime  time.Time `json:"create_time" db:"create_time"`
}

// ApiPostDetail 帖子详情接口的结构体
type ApiPostDetail struct {
	AuthorName       string                    `json:"author_name"`
	VoteNumber       int64                     `json:"vote_number"`
	*Post            `json:"post"`             //嵌入帖子结构体
	*CommunityDetail `json:"community_detail"` //嵌入社区信息
}
