package auth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt/v4"
)

// EncodeJwtToken 生成token
func EncodeFullAuthClaimJwtToken(userID, account, role string, secret []byte) (string, error) {
	c := FullAuthClaim{
		UserID:  userID,
		Account: account,
		Role:    role,
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
func DecodeFullAuthClaimJwtToken(tokenStr string, secret []byte) (*FullAuthClaim, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &FullAuthClaim{}, func(tk *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil {
		return nil, err
	}

	if claim, ok := token.Claims.(*FullAuthClaim); ok && token.Valid {
		return claim, nil
	}

	return nil, errors.New("invalid token ")
}
