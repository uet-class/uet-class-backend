package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func AbortWithError(c *gin.Context, code int, message interface{}) {
	fmt.Println(message)
	c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func ResponseHandler(c *gin.Context, code int, message interface{}) {
	fmt.Println(message)
	c.JSON(code, gin.H{"message": message})
}
