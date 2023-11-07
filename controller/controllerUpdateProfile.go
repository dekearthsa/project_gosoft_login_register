package controller

import (
	"context"
	"log"
	"net/http"
	"project_gosoft_login_register/model"

	"cloud.google.com/go/datastore"
	"github.com/gin-gonic/gin"
)

func ControllerUpdateProfile(c *gin.Context) {
	const PROJECTID = "nindocnx"
	const KIND = "UserProfile"
	var req model.Register
	if err := c.BindJSON(&req); err != nil {
		log.Println("err BindJSON =>", err)
		c.JSON(http.StatusBadRequest, gin.H{"status": "BadRequest"})
	}

	ctx := context.Background()
	client, err := datastore.NewClient(ctx, PROJECTID)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"status": "InternalServerError"})
	}

	var data []model.Register
	query := datastore.NewQuery(KIND).Filter("Username =", req.Username).Limit(1)
	if _, err := client.GetAll(ctx, query, &data); err != nil {
		log.Println(err)
	}

	payload := model.Register{
		Username:  data[0].Username,
		Password:  data[0].Password,
		Sex:       req.Sex,
		Age:       req.Age,
		Height:    req.Height,
		Weight:    req.Weight,
		Excercise: req.Excercise,
		Target:    req.Target,
		Meal:      req.Meal,
		TargetCal: req.TargetCal,
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

	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
