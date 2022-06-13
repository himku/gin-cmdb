package utils

import (
	"fmt"
	"github.com/dgrijalva/jwt-go/v4"
	"testing"
	"time"
)

/**
 * @Description
 * @Author sjie
 * @Date 2022/6/13 16:47
 **/

func TestGenToken(t *testing.T) {
	tokenClams := JwtCustomClaims{
		Username: "sjie",
		Password: "sjie",
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: jwt.At(time.Now().Add(time.Minute * 60)),
		},
	}
	token, err := MakeClamsToken(tokenClams)
	t.Log(token, err)
	orgToken, err := ParseClamsToken(token)
	t.Logf("%+v,%+v", orgToken, err)
}

func TestParseToken(t *testing.T) {
	tokenClams := JwtCustomClaims{Username: "sjie", Password: "sjie", StandardClaims: jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(time.Second * 60)),
	},
	}
	token, _ := MakeClamsToken(tokenClams)
	fmt.Println(token)
	claims, _ := ParseClamsToken(token)
	fmt.Println(claims)
	fmt.Println(claims.ExpiresAt)

}

func TestRefreshToken(t *testing.T) {
	tokenClams := JwtCustomClaims{Username: "sjie", Password: "sjie", StandardClaims: jwt.StandardClaims{
		ExpiresAt: jwt.At(time.Now().Add(time.Second * 60)),
	}}
	token, _ := MakeClamsToken(tokenClams)
	newToken, err := RefreshToken(token)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(newToken)
	claims, _ := ParseClamsToken(token)
	fmt.Println(claims)
}
