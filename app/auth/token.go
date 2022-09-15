package auth

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

func CreateToken(userID int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).UTC().Unix()
	claims["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}
