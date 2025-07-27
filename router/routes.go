package router

import "github.com/gin-gonic/gin"

func InitializeRoutes(router *gin.Engine) {
	v1 := router.Group("v1")
	{
		v1.GET("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.POST("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.DELETE("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.PUT("/opening", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
		v1.GET("/openings", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}
}