package routes

import (
	"go-api/controllers"
	"go-api/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/login", controllers.Login)
	r.POST("/register", controllers.Register)
	r.GET("/user", middleware.AuthMiddleware(), controllers.GetUser)

	return r
}
