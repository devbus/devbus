package daos

import (
	"github.com/devbus/devbus/common"
	"github.com/devbus/devbus/models"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"strings"
)

type UserDao struct {
	*gorm.DB
}

func NewUserDao(db *gorm.DB) *UserDao {
	return &UserDao{db}
}

func (dao *UserDao) GetUserByEmail(email string) *models.User {
	user := &models.User{}
	if err := dao.Where("email = ?", email).First(user).Error; err != nil {
		log.Debugf("failed to query user email: %v", err)
		return nil
	}
	return user
}

func (dao *UserDao) AddUser(user *models.User) error {
	if err := dao.Create(user).Error; err != nil {
		log.Debugf("%+v", errors.Wrapf(err, "failed to create new user, user: %+v", user))
		errStr := err.Error()
		errCode := common.ErrUnknown
		if strings.Contains(errStr, "devbus_user_name_key") {
			errCode = common.ErrConflictUserName
		} else if strings.Contains(errStr, "devbus_user_email_key") {
			errCode = common.ErrConflictEmail
		}
		return common.OpError{errCode}
	}
	return nil
}
