package services

import (
	"crypto/sha256"

	"encoding/hex"

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
	if email == "" || password == "" {
		return false
	}
	db := models.DBOpen()
	defer db.Close()
	dao := daos.NewUserDao(db)
	user := dao.GetUserByEmail(email)
	if user != nil {
		sum := sha256.Sum224([]byte(password + user.Salt))
		hashedPassword := hex.EncodeToString(sum[:])
		if hashedPassword == user.Password {
			return true
		}
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
