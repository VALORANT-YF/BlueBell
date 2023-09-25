package mysql

import (
	"bluebell/models"
	"database/sql"
	"errors"

	"go.uber.org/zap"
)

func GetCommunityListDao() ([]*models.Community, error) {
	sqlStr := "select community_id , community_name from community"
	var data []*models.Community
	//执行sql
	err := db.Select(&data, sqlStr)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Warn("there is no community in db")
			err = nil
			return data, err
		}
	}
	return data, err
}

// GetCommunityDetailById 根据id查询社区详情
// GetCommunityDetailById 根据 ID 查询社区详情
func GetCommunityDetailById(id int64) (communityDetail *models.CommunityDetail, err error) {
	communityDetail = new(models.CommunityDetail)
	sqlStr := `select community_id, community_name, introduction, create_time from community where community_id = ?`
	err = db.Get(communityDetail, sqlStr, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			zap.L().Warn("there is no communityDetail in db")
			err = errors.New("无效的id")
		}
		zap.L().Error("query data failed", zap.Error(err))
	}
	return
}
