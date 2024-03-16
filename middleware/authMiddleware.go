package middleware

import (
	controller "github.com/Nazarii14/golang-auth-jwt/controllers"
	"github.com/Nazarii14/golang-auth-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers)
	incomingRoutes.POST("/users/:user_id", controller.GetUser)
}
