package utils

import (
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func GenerateJWT(username string, firstName string, lastName string, email string, accountType string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":     username,
		"first_name":   firstName,
		"last_name":    lastName,
		"email":        email,
		"account_type": accountType,
		"exp":          time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(jwtSecret))

	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}
