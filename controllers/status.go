package controllers

import (
	"github.com/gin-gonic/gin"
)

func AbortWithError(c *gin.Context, code int, message interface{}) {
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func ResponseHandler(c *gin.Context, code int, message interface{}) {
	c.JSON(code, gin.H{"message": message})
}
