package services

import (
	"bytes"
	"crypto/sha256"

	"github.com/devbus/devbus/daos"
	"github.com/devbus/devbus/models"
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
	db := models.DBOpen()
	defer db.Close()
	dao := daos.NewUserDao(db)
	user := dao.GetUserByEmail(email)
	if user != nil {
		hashedPassword := sha256.Sum256([]byte(password + user.Salt))
		if bytes.Compare(hashedPassword[:], []byte(user.Password)) == 0 {
			log.Debug("password error, user: %s", email)
			return true
		}
	}
	return false
}

func (*userServiceImpl) Register(user *models.User) (err error) {
	db := models.DBOpen()
	defer db.Close()
	dao := daos.NewUserDao(db)
	return dao.AddUser(user)
}
