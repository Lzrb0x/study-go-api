package handler

import "github.com/gin-gonic/gin"

func CreateOpeningHandler(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "endpoint de POST",
	})
}