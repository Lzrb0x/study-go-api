package handler

import "github.com/gin-gonic/gin"

func GetOpeningHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "endpoint de GET",
	})
}