package define

import "github.com/dgrijalva/jwt-go"

type UserClaim struct {
	Id       int
	Identity string
	Name     string
	jwt.StandardClaims
}

var JwtKey = "go-cloud-key"

var EmailCodeLength = 6
var EmailCodeExpireTime = 300

var ListPageSizeDefault = 20
