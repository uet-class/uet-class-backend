package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InternalServerErrorHandler(err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
}

func UnauthorizedErrorHandler(err error, c *gin.Context) {
	c.JSON(http.StatusUnauthorized, gin.H{"error": err})
}

func MessageHandler(message interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": message})
}

func StatusHandler(status int, c *gin.Context) {
	c.JSON(status, gin.H{"status": status})
}
