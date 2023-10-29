package controller

import (
	"context"
	"log"
	"net/http"
	"project_gosoft_login_register/model"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func ControllerRegister(c *gin.Context) {

	const PROJECTID = "nindocnx"
	const KIND = "UserProfile"

	var req model.Register
	if err := c.BindJSON(&req); err != nil {
		log.Println("err BindJSON => ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad request"})
	}

	ctx := context.Background()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	payload := model.Register{
		Username:  req.Username,
		Password:  string(hashPassword),
		Sex:       req.Sex,
		Age:       req.Age,
		Height:    req.Height,
		Weight:    req.Weight,
		Excercise: req.Excercise,
		Target:    req.Target,
		Meal:      req.Meal,
		TargetCal: req.TargetCal,
	}

	// log.Println(payload)

	client, err := datastore.NewClient(ctx, PROJECTID)
	if err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't find projectID."})
	}

	key := datastore.IncompleteKey(KIND, nil)
	if _, err := client.Put(ctx, key, &payload); err != nil {
		log.Fatal(err)
		c.JSON(http.StatusInternalServerError, gin.H{"message": "can't insert data."})
	}

	c.JSON(http.StatusOK, gin.H{"status": "Create profile success."})

}
