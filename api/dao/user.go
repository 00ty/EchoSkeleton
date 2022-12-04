package dao

import (
	"EchoSkeleton/models"
	"EchoSkeleton/utils"

	"gorm.io/gorm"
)

type UserDao struct {
	db *gorm.DB
}

func NewUserDao() *UserDao {
	return &UserDao{db: utils.Db}
}

func (dao *UserDao) Find(uid int64) (*models.User, error) {
	user := models.User{}
	result := dao.db.First(&user, uid)
	if result.Error != nil {
		return nil, result.Error
	}
	return &user, nil
}
