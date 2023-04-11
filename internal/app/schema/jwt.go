package schema

import "github.com/golang-jwt/jwt"

// Claims jwt
type Claims struct {
	UID int // 用户 ID
	jwt.StandardClaims
}
