package models

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	mathrand "math/rand"
)

type User struct {
	ID             int32  `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	Password       string `json:"password"`
	FullName       string `json:"full_name"`
	IsActivate     bool   `json:"is_activate"`
	Salt           string
}

func (User) TableName() string {
	return "devbus_user"
}

func (u *User) HashPassword() {
	salt := make([]byte, 8)
	_, err := rand.Read(salt)
	if err != nil {
		log.Errorf("failed to genarate salt, error: %+v", err)
		if _, err := mathrand.Read(salt); err != nil {
			log.Errorf("failed to genarate salt using fallback method, error: %+v", err)
			u.Salt = "DEFAULT_SALT"
			return
		}
	}
	u.Salt = hex.EncodeToString(salt)
	sum := sha256.Sum224([]byte(u.Password + u.Salt))
	u.Password = hex.EncodeToString(sum[:])
}