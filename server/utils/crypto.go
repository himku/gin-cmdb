package utils

import (
	"golang.org/x/crypto/bcrypt"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/13 14:07
 **/

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
