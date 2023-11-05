package haddler

import (
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const KEY = "kdfoskovkoskodfkoasdlpslo"
const EXPTIME = 120

func HaddlerGenerateToken(Username string) (string, error) {
	var signKey = []byte(KEY)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["Username"] = Username
	claims["Exp"] = time.Now().Add(time.Minute * EXPTIME).Unix()
	claims["Create"] = time.Now().Add(time.Minute).Unix()
	tokenString, err := token.SignedString(signKey)
	if err != nil {
		log.Println("can't create json web token.")
		return "-", err
	}
	return tokenString, nil

}
