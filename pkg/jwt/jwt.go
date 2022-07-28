package jwt

import (
	"github.com/dgrijalva/jwt-go"
	"log"
	"time"
)

// secret key used to sign tokens
var (
	SecretKey = []byte("secret")
)

// returns a generated a JWT token and assigns username to its clams
func GenerateToken(username string) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	//Create a map to store our claims
	claims := token.Claims.(jwt.MapClaims)
	/* Set token claims username and exo */
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()

	tokenStr, err := token.SignedString(SecretKey)
	if err != nil {
		log.Fatal("Error in generating key")
		return "", err
	}
	return tokenStr, nil
}

// parses a JWT and returns the username in it claims
func ParseToken(tokenStr string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return SecretKey, nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		return username, nil
	} else {
		return "", err
	}
}
