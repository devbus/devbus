package services

import (
	"crypto/sha256"

	"encoding/hex"

	"github.com/devbus/devbus/common"
	"github.com/devbus/devbus/daos"
	"github.com/devbus/devbus/models"
)

type UserService interface {
	Auth(email, password string) (ok bool)
	Register(user *models.User) (err error)
	ModifyPassword(email, oldPassword, newPassword string) (err error)
}

func GetUserService() UserService {
	return &userServiceImpl{}
}

type userServiceImpl struct {
}

func (srv *userServiceImpl) Auth(email, password string) (ok bool) {
	if email == "" || password == "" {
		return false
	}
	db := models.DBOpen()
	defer db.Close()
	dao := daos.NewUserDao(db)
	user := dao.GetUserByEmail(email)
	if user != nil {
		return srv.doAuth(user, password)
	}
	log.Debugf("password error, user: %s", email)
	return false
}

func (*userServiceImpl) Register(user *models.User) (err error) {
	db := models.DBOpen()
	defer db.Close()
	user.HashPassword()
	dao := daos.NewUserDao(db)
	return dao.AddUser(user)
}

func (srv *userServiceImpl) ModifyPassword(email, oldPassword, newPassword string) (err error) {
	if email == "" || oldPassword == "" || newPassword == ""{
		return common.OpError{Code: common.ErrUnknown}
	}
	if oldPassword == newPassword {
		return common.OpError{Code: common.ErrUserSamePassword}
	}
	db := models.DBOpen()
	db = db.Begin()
	db.LogMode(true)
	defer db.Close()
	dao := daos.NewUserDao(db)
	user := dao.GetUserByEmail(email)
	if user == nil {
		return common.OpError{Code: common.ErrUserNotFound}
	}
	if !srv.doAuth(user, oldPassword) {
		return common.OpError{Code: common.ErrUserPasswrod}
	}
	sum := sha256.Sum224([]byte(newPassword + user.Salt))
	user.Password = hex.EncodeToString(sum[:])
	if err = dao.ModifyPassword(user); err != nil {
		db.Rollback()
		return
	}
	db.Commit()
	return
}

func (*userServiceImpl) doAuth(user *models.User, password string) bool {
	sum := sha256.Sum224([]byte(password + user.Salt))
	hashedPassword := hex.EncodeToString(sum[:])
	return hashedPassword == user.Password
}