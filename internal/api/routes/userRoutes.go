package routes

import (
	"github.com/Lzrb0x/study-go-api.git/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupUserRoutes(router *gin.Engine, userHandler *handlers.UserHandler) {
	userRoutes := router.Group("/users")
	{
		userRoutes.GET("/", userHandler.GetAllUsers)
	}
}
