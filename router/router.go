package router


import "github.com/gin-gonic/gin"

func SetupRouter() {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, "Pong")
	})

	r.Run(":8080")
}