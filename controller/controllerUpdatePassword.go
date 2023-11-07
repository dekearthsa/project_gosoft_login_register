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

type UserPasword struct {
	Username string
	Password string
}

func ControllerUpdatePassword(c *gin.Context) {
	const PROJECTID = "nindocnx"
	const KIND = "UserProfile"

	var req UserPasword
	if err := c.BindJSON(&req); err != nil {
		log.Println("err BindJSON => ", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "Bad request"})
	}
	// log.Println("req => ", req)

	ctx := context.Background()
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	client, err := datastore.NewClient(ctx, PROJECTID)

	var data []model.Register

	query := datastore.NewQuery(KIND).Filter("Username =", req.Username).Limit(1)
	if _, err := client.GetAll(ctx, query, &data); err != nil {
		log.Println(err)
	}

	data[0].Password = string(hashPassword)
	log.Println(data)

	payload := model.Register{
		Username:  data[0].Username,
		Password:  string(hashPassword),
		Age:       data[0].Age,
		Excercise: data[0].Excercise,
		Height:    data[0].Height,
		Meal:      data[0].Meal,
		Sex:       data[0].Sex,
		Target:    data[0].Target,
		TargetCal: data[0].TargetCal,
		Weight:    data[0].Weight,
	}

	keyEntity := datastore.NameKey(KIND, req.Username, nil)
	tx, err := client.NewTransaction(ctx)
	if err != nil {
		log.Println("client.NewTransaction => ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "NewTransaction error."})
	}

	if _, err := tx.Put(keyEntity, &payload); err != nil {
		log.Println("tx.Put => ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "put serror."})
	}
	if _, err := tx.Commit(); err != nil {
		log.Println("tx.Commit => ", err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "Commit error."})
	}

	c.JSON(http.StatusOK, gin.H{"status": "update success."})

}
