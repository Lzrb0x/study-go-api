package handler

import "github.com/gin-gonic/gin"

func GetOpeningHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "endpoint de GET",
	})
}

func CreateOpeningHandler(c *gin.Context) {
	c.JSON(201, gin.H{
		"message": "endpoint de POST",
	})
}

func DeleteOpeningHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "endpoint de DELETE",
	})
}

func UpdateOpeningHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "endpoint de PUT",
	})
}

func GetOpeningsHandler(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "endpoint de GET para múltiplas aberturas",
	})
}
