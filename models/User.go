package models

type User struct {
	ID   int64
	Uid  int64 `gorm:"primaryKey"`
	Gold int64
	Code string
}

// 自定义表名
func (User) TableName() string {
	return "bot_user1"
}
