package xjwt

import (
	"errors"
	"github.com/golang-jwt/jwt"
	"time"
)

var (
	TokenExpired     = errors.New("token is expired")
	TokenNotValidYet = errors.New("token not active yet")
	TokenMalformed   = errors.New("that's not even a token")
	TokenInvalid     = errors.New("couldn't handle this token")
)

type XJwt struct {
	// 签名秘钥
	secret []byte
	expire int64
	buffer int64
	isuser string
}

func NewXJwt(secret []byte, expire, buffer int64, isuser string) *XJwt {
	return &XJwt{secret: secret, expire: expire, buffer: buffer, isuser: isuser}
}

// CustomClaims
// Author [SliverFlow]
// @desc token 存储的信息 claims
type CustomClaims struct {
	jwt.StandardClaims
	UUID       string
	UserID     string
	BufferTime int64
}

// SendToken
// Author [SliverFlow]
// @desc 创建 token
// @param (userId, uuid string)
// @return (token string, err error)
func (x *XJwt) SendToken(userId, uuid string) (token string, err error) {
	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + x.expire,
			Issuer:    x.isuser,
			NotBefore: time.Now().Unix(),
		},
		UUID:       uuid,
		UserID:     userId,
		BufferTime: x.buffer,
	}
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return t.SignedString(x.secret)
}

// sendTokenWithClaims
// Author [SliverFlow]
// @desc 创建 token
// @param (userId, uuid string)
// @return (token string, err error)
func (x *XJwt) sendTokenWithClaims(custom *CustomClaims) (token string, err error) {
	t := jwt.NewWithClaims(jwt.SigningMethodHS256, custom)
	return t.SignedString(x.secret)
}

// ParseToken
// Author [SliverFlow]
// @desc 解析 token
// @param (token string)
// @return (userId, uuid string,err error)
func (x *XJwt) ParseToken(token string) (custom *CustomClaims, err error) {
	withClaims, err := jwt.ParseWithClaims(token, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return x.secret, nil
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
	if withClaims != nil {
		// 断言成功 而且 token 有效
		if custom, ok := withClaims.Claims.(*CustomClaims); ok && withClaims.Valid {
			return custom, nil
		}
		return nil, TokenInvalid
	} else {
		return nil, TokenInvalid
	}
}

// Renewal
// Author [SliverFlow]
// @desc token 续期
// @param (token string)
// @return (userId, uuid string,err error)
func (x *XJwt) Renewal(oldToken string) (token string, err error) {
	custom, err := x.ParseToken(oldToken)
	custom.ExpiresAt = time.Now().Unix() + x.expire // 更新过期时间
	if err != nil {
		return "", err
	}
	newToken, err := x.sendTokenWithClaims(custom)
	if err != nil {
		return "", nil
	}
	return newToken, nil
}
