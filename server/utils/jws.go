package utils

import (
	"errors"
	"github.com/dgrijalva/jwt-go/v4"
	"time"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/13 15:14
 **/

const Secret = "cmdb"

type JwtCustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func MakeClamsToken(obj JwtCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, obj)
	tokenString, err := token.SignedString([]byte(Secret))
	return tokenString, err
}

func ParseClamsToken(token string) (*JwtCustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*JwtCustomClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func RefreshToken(tokenStr string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenStr, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(Secret), nil
	})
	if err != nil {
		return "", err
	}
	if token != nil {
		if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
			jwt.TimeFunc = time.Now
			claims.StandardClaims.ExpiresAt = jwt.At(time.Now().Add(time.Minute * 60))
			return MakeClamsToken(JwtCustomClaims{})
		}
	}
	return "", errors.New("refresh Token Failed")
}
