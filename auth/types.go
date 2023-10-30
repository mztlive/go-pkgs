package auth

import (
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// SimpleAuthClaim
// 一个简单的实现了jwt.Claims接口的结构体， 只包含了用户ID
type SimpleAuthClaim struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

type FullAuthClaim struct {
	UserID  string `json:"userId"`
	Account string `json:"account"`
	Role    string `json:"role"`
	jwt.StandardClaims
}

// TokenExpireDuration token过期时间
const TokenExpireDuration = 120 * time.Hour
