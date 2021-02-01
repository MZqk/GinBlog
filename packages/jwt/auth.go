package jwt

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("This is secret!")

type Cliams struct {
	Username string
	Password string
	jwt.StandardClaims
}

func GenerateToken(username, password string) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)
	//expireTime := nowTime.Add(ini.ExpireTime * time.Hour)
	claims := Cliams{
		username,
		password,
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "ginblog",
		},
	}
	tokenCliams := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenCliams.SignedString(jwtSecret)
	fmt.Println("token", token, err)
	return token, err
}

func ParseToken(token string) (*Cliams, error) {
	tokenCliams, err := jwt.ParseWithClaims(token, &Cliams{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if tokenCliams != nil {
		if claims, ok := tokenCliams.Claims.(*Cliams); ok && tokenCliams.Valid {
			return claims, nil
		}
	}
	return nil, err
}
