package auth

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"os"
	"strconv"
	"strings"
	"time"
)

func CreateToken(userID int64) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["exp"] = time.Now().Add(time.Hour * 6).Unix()
	claims["userId"] = userID
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("TOKEN_SECRET")))
}

func ValidateToken(c *gin.Context) error {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, getVerifyKey)
	if err != nil {
		fmt.Println(err)
		return err
	}

	if _, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return nil
	}
	return fmt.Errorf("invalid token")
}

func ExtractUserID(c *gin.Context) (int64, error) {
	tokenString := extractToken(c)
	token, err := jwt.Parse(tokenString, getVerifyKey)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userID, err := strconv.ParseInt(fmt.Sprintf("%.0f", claims["userId"]), 10, 64)
		if err != nil {
			return 0, err
		}
		return userID, nil
	}

	return 0, fmt.Errorf("invalid token")
}
func extractToken(c *gin.Context) string {
	token := c.GetHeader("Authorization")
	if len(strings.Split(token, " ")) == 2 {
		return strings.Split(token, " ")[1]
	}
	return ""
}

func getVerifyKey(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected assign method")
	}
	return []byte(os.Getenv("TOKEN_SECRET")), nil
}
