package redis

import (
	"errors"
	"github.com/go-redis/redis"
	"math"
	"strconv"
	"time"
)

const (
	scorePerVote = 432           //投一票的分数
	maxPostTime  = 7 * 24 * 3600 //一周
)

func CreatePost(postId, communityId int64) error {
	pipeline := client.TxPipeline()
	pipeline.ZAdd(getRedisKey(KeyPostTimeZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	pipeline.ZAdd(getRedisKey(KeyPostScoreZSet), redis.Z{
		Score:  float64(time.Now().Unix()),
		Member: postId,
	})
	//把帖子id 加到社区的set
	cKey := getRedisKey(KeyCommunitySetPF + strconv.Itoa(int(communityId)))
	pipeline.SAdd(cKey, postId)
	_, err := pipeline.Exec()
	return err
}

func PostVote(userId, postId string, direction float64) error {
	//1.判断投票的限制 去redis中取帖子发布时间
	postTime := client.ZScore(getRedisKey(KeyPostTimeZSet), postId).Val()

	if float64(time.Now().Unix())-postTime > maxPostTime {
		return errors.New("超出有效投票时间")
	}
	//2.更新分数
	//首先查询之前的投票记录
	ov := client.ZScore(getRedisKey(KeyPostVotedZSet+postId), userId).Val()
	//如果这次投票的值和保存的一致,则提醒用户不允许重复投票
	if direction == ov {
		return errors.New("不允许重复投票")
	}
	var dir float64
	if direction > ov {
		dir = 1
	} else {
		dir = -1
	}
	diff := math.Abs(ov - direction) //计算两次投票的差值的绝对值
	pipeline := client.TxPipeline()
	pipeline.ZIncrBy(getRedisKey(KeyPostScoreZSet), dir*diff*scorePerVote, postId)
	//3.记录用户为该帖子投过票的数据
	if direction == 0 {
		pipeline.ZRem(getRedisKey(KeyPostVotedZSet+postId), userId)
	} else {
		pipeline.ZAdd(getRedisKey(KeyPostVotedZSet+postId), redis.Z{
			Score:  direction,
			Member: userId,
		})
	}
	_, err := pipeline.Exec()
	return err
}
