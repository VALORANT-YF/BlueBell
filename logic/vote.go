package logic

import (
	"bluebell/dao/redis"
	"bluebell/models"
	"go.uber.org/zap"
	"strconv"
)

//投票功能

// 投一票加432分 , 86400/200 -> 200张赞成票可以给帖子续上一天

//投票的限制:每个帖子自发表之日只允许有一周的投票时间,到期时候,将redis中保存的赞成票和反对票数存储到mysql中,并且删除 对应帖子的KeyPostVotedZSet

// PostVote 为帖子投票的函数
func PostVote(userId int64, p *models.VoteData) error {
	zap.L().Debug("VoteForPost",
		zap.Int64("userID", userId),
		zap.Int64("postID", p.PostId),
		zap.Int8("direction", p.Direction))
	return redis.PostVote(strconv.Itoa(int(userId)), strconv.Itoa(int(p.PostId)), float64(p.Direction))
}
