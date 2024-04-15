package config

import (
	"go.uber.org/zap"
	"golang.org/x/crypto/bcrypt"
)

func EncryptPwd(pwd string) (string, error) {
	// encrypt pwd
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(pwd), bcrypt.DefaultCost)
	if err != nil {
		logger.Error("error encrypting password", zap.String("error", err.Error()))
		return "", err
	}

	return string(hashedPwd), nil

}
