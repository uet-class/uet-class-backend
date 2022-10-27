package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func ErrorHandler(err error, c *gin.Context) {
	c.JSON(http.StatusInternalServerError, gin.H{"error": err})
}

func MessageHandler(obj interface{}, c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": obj})
}
