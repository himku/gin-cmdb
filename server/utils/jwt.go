package utils

import (
	"context"
	"errors"
	"fmt"
	"gin-cmdb/server/config"
	"github.com/dgrijalva/jwt-go/v4"
	"strconv"
	"time"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/13 15:14
 **/

var initConfig = config.NewConfig()

type JwtCustomClaims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func MakeClamsToken(obj JwtCustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, obj)
	tokenString, err := token.SignedString([]byte(initConfig.Jwt.SecretKey))
	return tokenString, err
}

func ParseClamsToken(token string) (*JwtCustomClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &JwtCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(initConfig.Jwt.SecretKey), nil
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
		return []byte(initConfig.Jwt.SecretKey), nil
	})
	if err != nil {
		return "", err
	}
	if token != nil {
		if claims, ok := token.Claims.(*JwtCustomClaims); ok && token.Valid {
			customClaims := JwtCustomClaims{Username: claims.Username, Password: claims.Password, StandardClaims: jwt.StandardClaims{ExpiresAt: jwt.At(time.Now().Add(time.Minute * 60))}}
			return MakeClamsToken(customClaims)
		}
	}
	return "", errors.New("refresh Token Failed")
}

func GetBlackListKey(tokenStr string) string {
	return "jwt_black_list:" + MD5([]byte(tokenStr))
}

// JoinBlackList token 加入黑名单
func JoinBlackList(token string) (err error) {
	nowUnix := time.Now().Unix()
	ExpireTime, _ := ParseClamsToken(token)
	timer := time.Duration(ExpireTime.ExpiresAt.Unix()-nowUnix) * time.Second
	// 将 token 剩余时间设置为缓存有效期，并将当前时间作为缓存 value 值
	err = InitializeRedis().SetNX(context.Background(), GetBlackListKey(token), nowUnix, timer).Err()
	fmt.Println(err)
	return
}

// IsInBlacklist token 是否在黑名单中
func IsInBlacklist(tokenStr string) bool {
	joinUnixStr, err := InitializeRedis().Get(context.Background(), GetBlackListKey(tokenStr)).Result()
	joinUnix, err := strconv.ParseInt(joinUnixStr, 10, 64)
	if joinUnixStr == "" || err != nil {
		return false
	}
	// JwtBlacklistGracePeriod 为黑名单宽限时间，避免并发请求失效
	if time.Now().Unix()-joinUnix < initConfig.JwtBlackList {
		return false
	}
	return true
}
