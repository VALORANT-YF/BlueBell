package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
)

// GetCommunityListService 查询所有的community并返回
func GetCommunityListService() ([]*models.Community, error) {
	data, err := mysql.GetCommunityListDao()
	if err != nil {
		return data, err
	}
	return data, nil
}

// GetCommunityDetailService 根据id获取社区详情
func GetCommunityDetailService(id int64) (*models.CommunityDetail, error) {
	communityDetails, err := mysql.GetCommunityDetailById(id)
	return communityDetails, err
}
