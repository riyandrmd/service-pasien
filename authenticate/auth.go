package authenticate

import (
	"github.com/dgrijalva/jwt-go"
)

var AuthSecret = "rumah-sakit"
var RefreshAuthSecret = "rumah-sakit-refresh"
var From = "http://localhost:8082"

type TokenDetail struct {
	AccessToken  string
	RefreshToken string
	AccessUuid   string
	RefreshUuid  string
	AtExpires    int64
}

type JWTClaims struct {
	jwt.StandardClaims
	UID   int      `json:"uid"`
	Group []string `json:"group"`
}
