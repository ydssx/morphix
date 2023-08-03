package jwt

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var (
	SecretKey      = []byte("123456")    // 秘钥
	ExpireDuration = time.Hour * 24 * 30 // 过期时间
)

type Claims struct {
	Uid      int64  `json:"uid"`
	Username string `json:"username"`
	Role     string `json:"role"`
	jwt.StandardClaims
}

func GenerateToken(uid int64, username, role string) (string, error) {
	claims := &Claims{
		Uid:      uid,
		Username: username,
		Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(ExpireDuration).Unix(),
			Issuer:    "my_app",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(SecretKey)
}

func VerifyToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*Claims)
	if !ok || !token.Valid {
		return nil, jwt.ErrInvalidKey
	}
	return claims, nil
}
