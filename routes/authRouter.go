package routes

import (
	"github.com/gin-gonic/gin"
	controller "github.com/harsh082ip/JWT-AUTHORIZATION-golang/controllers"
)

func AuthRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.POST("users/signup", controller.GetUser())
	incomingRoutes.POST("users/login", controller.Login())
}
