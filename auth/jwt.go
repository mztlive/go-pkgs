package auth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// SimpleAuthClaim
// 一个简单的实现了jwt.Claims接口的结构体， 只包含了用户ID
type SimpleAuthClaim struct {
	UserID string `json:"userId"`
	jwt.StandardClaims
}

// TokenExpireDuration token过期时间
const TokenExpireDuration = 120 * time.Hour

// EncodeJwtToken 生成token
func EncodeJwtToken(userID string, secret []byte) (string, error) {
	c := SimpleAuthClaim{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
			Issuer:    "TEST001",
		},
	}

	//使用指定的签名方法创建签名对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	// 使用指定的secret签名并获得完整的编码后的字符串token
	return token.SignedString(secret)
}

// DecodeJwtToken 解析token
func DecodeJwtToken(tokenStr string, secret []byte) (*SimpleAuthClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &SimpleAuthClaim{}, func(tk *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claim, ok := token.Claims.(*SimpleAuthClaim); ok && token.Valid {
		return claim, nil
	}

	return nil, errors.New("invalid token ")
}
