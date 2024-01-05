package utils

import (
	"errors"
	"github.com/golang-jwt/jwt/v4"
	"os"
	"time"
)

// 一些常量
var (
	TokenExpired     error = errors.New("token is expired")
	TokenNotValidYet error = errors.New("token not active yet")
	TokenMalformed   error = errors.New("that's not even a token")
	TokenInvalid     error = errors.New("couldn't handle this token")
)

// CustomClaims 载荷，可以加一些自己需要的信息
type CustomClaims struct {
	UID int
	jwt.RegisteredClaims
}

// JWT 签名结构
type JWT struct {
	SigningKey []byte `json:"signing_key"`
}

// NewJWT 新建一个jwt实例
func NewJWT() *JWT {
	return &JWT{
		SigningKey:[]byte(os.Getenv("JWT_SIGNING_KEY")),
	}
}

// createToken 生成一个token
func (j *JWT) createToken(claims CustomClaims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.SigningKey)
}

// GenerateToken 生成令牌
func GenerateToken( uId int) string {
	j := NewJWT()
	type cus struct {
		UID int
		jwt.RegisteredClaims
	}
	claims := cus{
		uId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}

	token, err :=  j.createToken(CustomClaims(claims))
	if err != nil {

		return err.Error()
	}
	//log.Println("--->生成的token-->：" + token)
	return token
}

// RefreshToken 更新token
func (j *JWT) RefreshToken(tokenString string) (string, error) {
	jwt.TimeFunc = func() time.Time {
		return time.Unix(0, 0)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		return "", err
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		jwt.TimeFunc = time.Now
		claims.ExpiresAt =jwt.NewNumericDate(time.Now().Add(time.Hour))
		return j.createToken(*claims)
	}
	return "", TokenInvalid
}

// ParseToken 解析 Tokne
func (j *JWT) ParseToken(tokenString string) (*CustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return j.SigningKey, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, TokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				// Token is expired
				return nil, TokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, TokenNotValidYet
			} else {
				return nil, TokenInvalid
			}
		}
	}
	if token == nil {
		return nil, TokenInvalid
	}
	if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, TokenInvalid
}

