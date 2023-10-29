package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ControllerDebug(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}
