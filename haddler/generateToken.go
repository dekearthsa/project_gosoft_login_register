package haddler

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt"
)

const KEY = "kdfoskovkoskodfkoasdlpslo"

func HaddlerGenerateToken(Username string) (string, error) {
	var signKey = []byte(KEY)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["Username"] = Username
	claims["exp"] = time.Now().Add(time.Minute * 120).Unix()
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		log.Println("can't create json web token.")
		return "-", err
	}
	return tokenString, nil

}
