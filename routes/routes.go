package routes

import (
	"y-api/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	userRoutes := r.Group("/")
	{
		userRoutes.POST("/register", controllers.RegisterUser)

		// NOT IMPLEMENTED
		// userRoutes.POST("/login", controllers.LoginUser)
		// userRoutes.POST("/logout", controllers.LogoutUser)
	}

	return r
}
