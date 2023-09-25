package models

import "time"

// Community 查询community_id 和 community_name 字段
type Community struct {
	CommunityId   int64  `json:"community_id" db:"community_id"`
	CommunityName string `json:"community_name" db:"community_name"`
}

// CommunityDetail 社区详情
type CommunityDetail struct {
	Id            int64     `json:"id" db:"id"`
	CommunityId   int64     `json:"community_id" db:"community_id"`
	CommunityName string    `json:"community_name" db:"community_name"`
	Introduction  string    `json:"introduction,omitempty" db:"introduction"`
	CreateTime    time.Time `json:"create_time" db:"create_time"`
	UpdateTime    time.Time `json:"update_time" db:"update_time"`
}
