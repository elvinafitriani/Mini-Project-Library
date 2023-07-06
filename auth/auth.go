package auth

import (
	"github.com/dgrijalva/jwt-go"
)

type TokenDetail struct {
	AccsesToken string
	RefAccToken string
	ExpAt       int64
	AccId       string
	RefAccID    string
}

type Claims struct {
	jwt.StandardClaims
	Uid  int
	User string
}