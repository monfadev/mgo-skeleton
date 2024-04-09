package helpers

import (
	"fmt"
	"mgo-skeleton/bin/modules/auth/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var key = []byte("mgoskeleton")

type JWTCustomClaims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

func GenerateToken(user *models.UserModel) (string, error) {
	fmt.Printf("GenerateToken struct with pointer is: %v", user.ID)
	claims := JWTCustomClaims{
		user.ID,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(120 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(key)

	return ss, err
}
