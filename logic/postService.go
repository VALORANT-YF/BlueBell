package logic

import (
	"bluebell/dao/mysql"
	"bluebell/dao/redis"
	"bluebell/models"
	"bluebell/pkg/snowflake"
	"fmt"
	"go.uber.org/zap"
)

func CreatePost(post *models.Post) (err error) {
	//生成post_id
	post.PostId = snowflake.GenID()
	//将帖子信息保存到数据库
	err = mysql.CreatePostInsert(post)
	if err != nil {
		return err
	}
	//将创建时间保存到redis 数据库
	err = redis.CreatePost(post.PostId, post.CommunityId)
	return
}

// GetPostDetailById 根据帖子id获取帖子详情
func GetPostDetailById(postId int64) (data *models.ApiPostDetail, err error) {
	//查询并且组合接口需要的数据
	//首先查询帖子详情
	postData, err := mysql.SelectPostDetailById(postId)
	if err != nil {
		zap.L().Error("mysql.GetPostById failed", zap.Error(err))
		return
	}
	//根据作者id 查询发帖人
	user, err := mysql.SelectPostAuthorById(postData.AuthorID)
	if err != nil {
		zap.L().Error("mysql.SelectPostAuthorById failed", zap.Int64("pid", postId), zap.Error(err))
		return
	}
	//根据社区id 查询社区的详细信息
	community, err := mysql.GetCommunityDetailById(postData.CommunityId)
	if err != nil {
		zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("pid", postId), zap.Error(err))
		return
	}
	data = &models.ApiPostDetail{
		AuthorName:      user,
		Post:            postData,
		CommunityDetail: community,
	}
	fmt.Println("@@@@@")
	return
}

// GetPostList 获取帖子列表
func GetPostList(postList *models.PostList) (data []*models.ApiPostDetail, err error) {
	posts, err := mysql.SelectPostList(postList)
	if err != nil {
		return nil, err
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	for _, post := range posts {
		//根据作者id 查询发帖人
		user, err := mysql.SelectPostAuthorById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.SelectPostAuthorById failed", zap.Int64("pid", post.AuthorID), zap.Error(err))
			continue
		}
		//根据社区id 查询社区的详细信息
		community, err := mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("pid", post.CommunityId), zap.Error(err))
			continue
		}
		postDetail := &models.ApiPostDetail{
			AuthorName:      user,
			Post:            post,
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetPostList2(p *models.ParamPostList) (data []*models.ApiPostDetail, err error) {
	//在redis 中查询id列表
	ids, err := redis.GetPostIdsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		return
	}
	zap.L().Debug("ids", zap.Any("ids", ids))
	//根据id 去mysql数据库去查询帖子详情
	posts, err := mysql.GetPostListByIds(ids)
	zap.L().Debug("posts", zap.Any("posts", posts))
	if err != nil {
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	//提前查询每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 将帖子的作者和分区 查询出来 填充到帖子中
	for idx, post := range posts {
		//根据作者id 查询发帖人
		user, err := mysql.SelectPostAuthorById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.SelectPostAuthorById failed", zap.Int64("pid", post.AuthorID), zap.Error(err))
			continue
		}
		zap.L().Debug("user is nil", zap.Any("user", user))
		//根据社区id 查询社区的详细信息
		community, err := mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("pid", post.CommunityId), zap.Error(err))
			continue
		}
		zap.L().Debug("user is nil", zap.Any("community", community))
		postDetail := &models.ApiPostDetail{
			AuthorName:      user,
			Post:            post,
			VoteNumber:      voteData[idx],
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}

func GetCommunityPostList2(p *models.ParamCommunityPostList) (data []*models.ApiPostDetail, err error) {
	//在redis 中查询id列表

	ids, err := redis.GetCommunityPostIdsInOrder(p)
	if err != nil {
		return
	}
	if len(ids) == 0 {
		return
	}
	zap.L().Debug("ids", zap.Any("ids", ids))
	//根据id 去mysql数据库去查询帖子详情
	posts, err := mysql.GetPostListByIds(ids)
	zap.L().Debug("posts", zap.Any("posts", posts))
	if err != nil {
		return
	}
	data = make([]*models.ApiPostDetail, 0, len(posts))
	//提前查询每篇帖子的投票数
	voteData, err := redis.GetPostVoteData(ids)
	if err != nil {
		return
	}

	// 将帖子的作者和分区 查询出来 填充到帖子中
	for idx, post := range posts {
		//根据作者id 查询发帖人
		user, err := mysql.SelectPostAuthorById(post.AuthorID)
		if err != nil {
			zap.L().Error("mysql.SelectPostAuthorById failed", zap.Int64("pid", post.AuthorID), zap.Error(err))
			continue
		}
		zap.L().Debug("user is nil", zap.Any("user", user))
		//根据社区id 查询社区的详细信息
		community, err := mysql.GetCommunityDetailById(post.CommunityId)
		if err != nil {
			zap.L().Error("mysql.GetCommunityDetailById failed", zap.Int64("pid", post.CommunityId), zap.Error(err))
			continue
		}
		zap.L().Debug("user is nil", zap.Any("community", community))
		postDetail := &models.ApiPostDetail{
			AuthorName:      user,
			Post:            post,
			VoteNumber:      voteData[idx],
			CommunityDetail: community,
		}
		data = append(data, postDetail)
	}
	return
}
