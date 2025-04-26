package auth

import (
	"github.com/khemingkapat/been_chillin/objects"
	"golang.org/x/crypto/bcrypt"
)

func EncryptUser(user *object.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hashedPassword)
	return nil
}
