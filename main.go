package main

import (
	"log"
	"project_gosoft_login_register/controller"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const PORT string = ":8111"

func main() {
	router := gin.New()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"POST", "PUT", "PATCH", "DELETE", "GET"},
		AllowHeaders: []string{"Content-Type,access-control-allow-origin,access-control-allow-headers,Authorization"},
	}))

	router.GET("/api/debug", controller.ControllerDebug)
	router.POST("/api/register", controller.ControllerRegister)
	router.POST("/api/login", controller.ControllerLogin)
	router.GET("/api/check/auth", controller.ControllerCheckToken)
	router.POST("/api/update/password", controller.ControllerUpdatePassword)
	router.POST("/api/update/profile", controller.ControllerUpdateProfile)

	err := router.Run(PORT)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("service run find at port", PORT)
}
