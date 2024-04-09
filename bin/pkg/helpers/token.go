package helpers

import (
	"errors"
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
	fmt.Println("\n")
	fmt.Printf("GenerateToken struct with pointer is: %v", user.ID)
	fmt.Println("\n")
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

func ValidateToken(tokenStr string) (*int, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &JWTCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return key, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid access")
		}
		return nil, errors.New("access expired")
	}

	/// casting to struct
	claims, ok := token.Claims.(*JWTCustomClaims)
	if !ok || !token.Valid {
		return nil, errors.New("access invalid and expired")
	}

	return &claims.ID, nil
}
