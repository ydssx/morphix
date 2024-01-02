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

// GenerateToken 生成一个包含uid、用户名、角色信息的JWT token。
// uid是用户ID。username是用户名。role是用户角色。
// 它会设置token的过期时间为ExpireDuration。签名方法为HS256。
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

// VerifyToken 验证JWT token是否有效。
// tokenString是要验证的token字符串。
// 它会解析token,验证签名和过期时间。
// 如果验证成功,返回Claims信息,否则返回错误。
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
