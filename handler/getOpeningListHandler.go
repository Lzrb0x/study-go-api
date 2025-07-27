package handler

import "github.com/gin-gonic/gin"

func GetOpeningsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "endpoint de GET para múltiplas aberturas",
	})
}