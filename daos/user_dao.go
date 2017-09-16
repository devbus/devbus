package daos

import (
	"github.com/devbus/devbus/common"
	"github.com/devbus/devbus/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) GetUserByEmail(email string) *models.User {
	user := &models.User{Email: email}
	if dao.DB.Find(user); dao.Error != nil {
		log.Debug("failed to query user email: %v", dao.Error)
		return nil
	}
	return user
}

func (dao *UserDao) AddUser(user *models.User) error {
	if !dao.DB.NewRecord(user) {
		log.Debug("%+v", errors.Wrapf(dao.DB.Error, "failed to create new user, user: %+v", user))
		return common.OpError{}
	}
	return nil
}