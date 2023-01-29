package common

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"strconv"
)

type JWTClaims struct {
	UserId   int64  `json:"userid"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var (
	Secret = []byte("douyin")
	// TokenExpireDuration = time.Hour * 2 过期时间
)

// 生成token
func GenToken(userId int64, userName string) (string, error) {
	claims := JWTClaims{
		UserId:   userId,
		Username: userName,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "server",
			//ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),可用于设定token过期时间
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte("douyin"))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// 解析token
func ParsenToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// 检查token
func CheckToken(uid string, token string) error {
	Claims, err := ParsenToken(token)
	if err != nil {
		return err
	}
	if strconv.Itoa(int(Claims.UserId)) != uid {
		return err
	}
	return nil
}

// 账号密码参数检验
func Check_AccountParam(username string, password string) bool {
	//防sql注入
	//账户长度不得大于32位,密码长度不得大于32位
	if len(username) > 32 || len(password) > 32 {
		return false
	}
	return true
}

func ProcessUid(str string, uid string) string {
	return fmt.Sprintf("%d%d", str, uid)
}
