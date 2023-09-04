package xjwt

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
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
	secret          []byte
	expire          int64
	buffer          int64
	isuser          string
	blackListPrefix string
}

func NewXJwt(secret []byte, expire, buffer int64, isuser, blackListPrefix string) *XJwt {
	return &XJwt{secret: secret, expire: expire, buffer: buffer, isuser: isuser, blackListPrefix: blackListPrefix}
}

// GetUserUUID
// Author [SliverFlow]
// @desc 获取 userID
// @param (userId, uuid string)
// @return (token string, err error)
func GetUserUUID(ctx context.Context) (string, error) {
	claims, ok := ctx.Value("claims").(*CustomClaims)
	if !ok {
		return "", errors.New("userID 获取失败")
	}
	return claims.UUID, nil
}

// CustomClaims
// Author [SliverFlow]
// @desc token 存储的信息 claims
type CustomClaims struct {
	jwt.StandardClaims
	UUID       string
	BufferTime int64
}

// SendToken
// Author [SliverFlow]
// @desc 创建 token
// @param (userId, uuid string)
// @return (token string, err error)
func (x *XJwt) SendToken(uuid string) (token string, err error) {
	claims := &CustomClaims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + x.expire,
			Issuer:    x.isuser,
			NotBefore: time.Now().Unix(),
		},
		UUID:       uuid,
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

// IsRedisBlackList
// Author [SliverFlow]
// @desc 判断当前 token 是否在黑名单中
// @param token string, client *redis.Client
// @return bool
func (x *XJwt) IsRedisBlackList(token string, client *redis.Client) bool {
	ctx := context.Background()
	_, err := client.Get(ctx, x.blackListPrefix+token).Result()
	if err != nil {
		return false
	}
	return true
}

// RedisBlackList
// Author [SliverFlow]
// @desc 判断当前 token 是否在黑名单中
// @param (token string, client *redis.Client)
// @return error
func (x *XJwt) RedisBlackList(token string, client *redis.Client) error {
	ctx := context.Background()
	err := client.Set(ctx, x.blackListPrefix+token, "black lsit ", time.Duration(x.buffer)*time.Second).Err()
	if err != nil {
		return err
	}
	return nil
}
