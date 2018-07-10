package utils

import (
	"github.com/msleevi/com.msleevi/internal/routes/users/formats"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// GeneratePassword creates a Password
func GeneratePassword(plainText string) (ret string, err error) {
	result, err := bcrypt.GenerateFromPassword([]byte(plainText), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Unable to create password")
		return "", err
	}
	ret = string(result)
	return ret, nil
}

func ComparePassword(plainText string, password formats.Password) (ret bool, err error) {
	err = bcrypt.CompareHashAndPassword([]byte(password.Password), []byte(plainText))
	if err != nil {
		return false, err
	}
	return true, nil
}
