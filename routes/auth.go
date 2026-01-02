
package routes

import (
	"backend/controllers"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func AuthRoutes(r *gin.Engine) {

	api := r.Group("/api")
	{
		api.GET("/hello", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "Hello"})
		})

		api.POST("/register", controllers.Register)
		api.POST("/login", controllers.Login)

		api.GET("/profile", middleware.JWTAuth(), func(c *gin.Context) {
			c.JSON(200, gin.H{"user": c.MustGet("user")})
		})
	}
}
