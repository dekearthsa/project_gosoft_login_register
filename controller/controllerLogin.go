package controller

import (
	"context"
	"log"
	"net/http"
	"project_gosoft_login_register/haddler"
	"project_gosoft_login_register/model"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ControllerLogin(c *gin.Context) {
	const PROJECTID = "nindocnx"
	const KIND = "UserProfile"

	var req model.Login

	if err := c.BindJSON(&req); err != nil {
		log.Println("err BindJSON => ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad request"})
	}

	ctx := context.Background()

	client, err := datastore.NewClient(ctx, PROJECTID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't find projectID."})
	}
	defer client.Close()

	var data []model.Register
	q := datastore.NewQuery(KIND).Filter("Username =", req.Username).Limit(1)
	if _, err := client.GetAll(ctx, q, &data); err != nil {
		log.Println(err)
	}

	if len(data) == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
	} else {
		err := bcrypt.CompareHashAndPassword([]byte(data[0].Password), []byte(req.Password))
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
		}

		token, err := haddler.HaddlerGenerateToken(req.Username)
		c.JSON(http.StatusOK, gin.H{"token": token})
	}

}
