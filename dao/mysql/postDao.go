package mysql

import (
	"bluebell/models"
	"github.com/jmoiron/sqlx"
	"strings"
)

// CreatePostInsert 将帖子信息添加到数据库
func CreatePostInsert(postInformation *models.Post) (err error) {
	sqlStr := `insert into post (post_id , title , content , author_id , community_id) values (?,?,?,?,?)`
	//执行插入数据库的sql语句
	_, err = db.Exec(sqlStr, postInformation.PostId, postInformation.Title, postInformation.Content, postInformation.AuthorID, postInformation.CommunityId)
	return
}

// SelectPostDetailById 根据帖子id查询帖子详情
func SelectPostDetailById(postId int64) (postData *models.Post, err error) {
	post := new(models.Post)
	sqlStr := "select post_id , title , content , author_id , community_id , create_time from post where post_id = ?"
	//执行sql语句
	err = db.Get(post, sqlStr, postId)
	return post, err
}

// SelectPostList 获取帖子分页详情
func SelectPostList(postList *models.PostList) (posts []*models.Post, err error) {
	sqlStr := `select post_id , title , content , author_id , community_id , create_time from post ORDER BY create_time desc  limit ? , ?`
	posts = make([]*models.Post, 0, 2) //使用make 初始化切片 , 容量0 , 长度2
	//执行sql

	err = db.Select(&posts, sqlStr, (postList.Offset-1)*postList.Limit, postList.Limit)
	return
}

// GetPostListByIds 根据id查询帖子详情数据
func GetPostListByIds(ids []string) (postList []*models.Post, err error) {
	sqlStr := `select post_id , title , content , author_id , community_id , create_time from post where post_id in (?) order by FIND_IN_SET(post_id , ?)`
	query, args, err := sqlx.In(sqlStr, ids, strings.Join(ids, ","))
	if err != nil {
		return nil, err
	}
	query = db.Rebind(query)
	err = db.Select(&postList, query, args...)
	return
}
