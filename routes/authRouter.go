package routes

import (
	controller "github.com/Nazarii14/golang-auth-jwt/controllers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(incomingRoutes *gin.Engine) {
	incomingRoutes.POST("/auth/register", controller.Signup)
	incomingRoutes.POST("/auth/login", controller.Login)
	incomingRoutes.POST("/auth/logout", controller.Logout)
	incomingRoutes.POST("/auth/refresh-token", controller.RefreshToken)
}
