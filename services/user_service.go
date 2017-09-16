package services

import (
	"github.com/devbus/devbus/models"
	"github.com/devbus/devbus/daos"
	"crypto/sha256"
	"bytes"
	"github.com/astaxie/beego/logs"
	"github.com/astaxie/beego/orm"
)

type UserService interface {
	Auth(email, password string) (ok bool)
	Register(user *models.User) (err error)
}

func GetUserService() UserService {
	return &userServiceImpl{}
}

type userServiceImpl struct {

}

func (*userServiceImpl) Auth(email, password string) (ok bool) {
	o := orm.NewOrm()
	dao := daos.NewUserDao(o)
	user := dao.GetUserByEmail(email)
	if user != nil {
		hashedPassword := sha256.Sum256([]byte(password + user.Salt))
		if bytes.Compare(hashedPassword[:], []byte(user.Password)) == 0 {
			logs.Debug("password error, user: %s", email)
			return true
		}
	}
	return false
}

func (*userServiceImpl) Register(user *models.User) (err error) {
	o := orm.NewOrm()
	dao := daos.NewUserDao(o)
	return dao.AddUser(user)
}
