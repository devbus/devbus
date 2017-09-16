package daos

import (
	"github.com/astaxie/beego/orm"
	"github.com/devbus/devbus/models"
	"github.com/astaxie/beego/logs"
	"github.com/pkg/errors"
	"github.com/devbus/devbus/common"
)

type UserDao struct {
	orm orm.Ormer
}

func NewUserDao(orm orm.Ormer) *UserDao {
	return &UserDao{orm}
}

func (dao *UserDao) GetUserByEmail(email string) *models.User {
	user := &models.User{Email: email}
	if err := dao.orm.Read(user); err != nil {
		logs.Debug("failed to query user email: %v", err)
		return nil
	}
	return user
}

func (dao *UserDao) AddUser(user *models.User) error {
	if _, err := dao.orm.Insert(user); err != nil {
		logs.Debug("%+v", errors.Wrapf(err, "failed to create new user, user: %v", user))
		return common.OpError{}
	}
	return nil
}