package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func AuthenticateFailed(c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"status": "Unauthorized"})
}
