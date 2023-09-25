package redis

//redis key 存放常用的Key
//redis key 尽量使用命名空间方式来区分

const (
	KeyPrefix        = "bluebell:"
	KeyPostTimeZSet  = "post:time"  //zset;帖子及发帖时间
	KeyPostScoreZSet = "post:score" //zset;帖子及投票分数
	KeyPostVotedZSet = "post:voted" //zset;记录用户以及投票的类型

	KeyCommunitySetPF = "community:" //set;保存每个分区下帖子的id
)

// getRedisKey 给redis的key加上前缀
func getRedisKey(key string) string {
	return KeyPrefix + key
}
