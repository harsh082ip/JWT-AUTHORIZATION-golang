package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/harsh082ip/JWT-AUTHORIZATION-golang/controllers"
	"github.com/harsh082ip/JWT-AUTHORIZATION-golang/middleware"
)

func UserRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.Use(middleware.Authenticate())
	incomingRoutes.GET("/users", controller.GetUsers())
	incomingRoutes.GET("/users/{user_id}", controller.GetUser())
}
