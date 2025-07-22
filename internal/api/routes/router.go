package routes

import (
	"github.com/Lzrb0x/study-go-api.git/internal/api/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes() {

	router := gin.Default()

	userHandler := handlers.NewUserHandler()

	SetupUserRoutes(router, userHandler)

	router.Run(":8080")
}
