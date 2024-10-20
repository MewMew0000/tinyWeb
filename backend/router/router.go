package router

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	auth := r.Group("/api/auth")
	auth.POST("/login", controllers.Login)
	auth.POST("/register", controllers.Register)
	return r
}
