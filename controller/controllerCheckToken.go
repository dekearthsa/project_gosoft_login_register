package controller

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

const KEY = "kdfoskovkoskodfkoasdlpslo"

type Request struct {
	Username string
	Create   int64
	Exp      int64
}

func ControllerCheckToken(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")

	// log.Println("myToken => ", tokenString)
	if tokenString == "" {
		c.JSON(http.StatusUnauthorized, gin.H{"login": false, "desc": "Unauthorized"})
	} else {
		// token, err := jwt.Parse(tokenString, nil)
		// if err != nil {
		// 	log.Print(err)
		// }

		// claims, _ := token.Claims.(jwt.MapClaims)

		arrayToken := strings.Split(tokenString, ".")
		// log.Print(arrayToken[1])
		decoded, err := base64.RawURLEncoding.DecodeString((arrayToken[1]))

		if err != nil {
			log.Println(err)
		}

		data := Request{}
		json.Unmarshal([]byte(decoded), &data)
		// log.Println(data.Exp)
		timing := time.Now().Add(time.Minute).Unix()

		if data.Exp < timing {
			c.JSON(http.StatusUnauthorized, gin.H{"login": false, "desc": "session timeout."})
		} else {
			c.JSON(http.StatusOK, gin.H{"login": true, "desc": "ok"})
		}

	}
}
