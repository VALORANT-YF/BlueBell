package redis

import (
	"bluebell/models"
	"github.com/go-redis/redis"
	"strconv"
	"time"
)

func getIDSFormKey(key string, page, size int64) ([]string, error) {
	//确定查询的起始索引点
	start := (page - 1) * size
	end := start + size - 1
	// ZREVRANGE 按分数从大到小查询
	return client.ZRevRange(key, start, end).Result()
}

func GetPostIdsInOrder(p *models.ParamPostList) ([]string, error) {
	//从redis 中获取id
	//根据用户请求中携带的order 确定 需要的key
	key := getRedisKey(KeyPostTimeZSet)
	if p.Order == "score" {
		key = getRedisKey(KeyPostScoreZSet)
	}
	return getIDSFormKey(key, p.Page, p.Size)
}

// GetPostVoteData 根据ids 查询每篇帖子投赞同票的数据
func GetPostVoteData(ids []string) (data []int64, err error) {
	//data = make([]int64, 0, len(ids))
	//for _, id := range ids {
	//	key := getRedisKey(KeyPostVotedZSet + id)
	//	v1 := client.ZCount(key, "1", "1").Val() //统计每篇帖子的赞成票的数量
	//	data = append(data, v1)
	//}
	keys := make([]string, 0, len(ids))
	pipeline := client.Pipeline()
	for _, id := range ids {
		key := getRedisKey(KeyPostVotedZSet + id)
		keys = append(keys, key)
		pipeline.ZCount(key, "1", "1")
	}
	cmders, err := pipeline.Exec()
	if err != nil {
		return nil, err
	}
	for _, cmder := range cmders {
		v := cmder.(*redis.IntCmd).Val()
		data = append(data, v)
	}
	return
}

// GetCommunityPostIdsInOrder 按照社区查找ids
func GetCommunityPostIdsInOrder(p *models.ParamCommunityPostList) ([]string, error) {
	orderKey := getRedisKey(KeyPostTimeZSet)
	if p.Order == "score" {
		orderKey = getRedisKey(KeyPostScoreZSet)
	}
	//使用 zinterstore 把分区的帖子set 与 帖子分数的 zset 生成一个新的zset

	//针对新的zset 按之前的逻辑取数据

	//社区的key
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(p.CommunityId)))
	key := orderKey + strconv.Itoa(int(p.CommunityId))
	if client.Exists(orderKey).Val() < 1 {
		//不存在, 需要计算
		pipeline := client.Pipeline()
		pipeline.ZInterStore(key, redis.ZStore{
			Aggregate: "MAX",
		}, cKey, orderKey) //计算
		pipeline.Expire(key, 60*time.Second) //设置超时时间
		_, err := pipeline.Exec()
		if err != nil {
			return nil, err
		}
	}
	//根据key 查询ids
	return getIDSFormKey(key, p.Page, p.Size)
}
