package models

// User 用户注册时向表中添加的字段
type User struct {
	UserId   int64  `json:"user_id,string" db:"user_id"`
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}

// UserId 用户的id 和 user_id
type UserId struct {
	Id     int64 `json:"id" db:"id"`
	UserId int64 `json:"user_id" db:"user_id"`
}

func (u UserId) TableName() string {
	return "user"
}
