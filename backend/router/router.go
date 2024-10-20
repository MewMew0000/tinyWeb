package router

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(cors.New(cors.Config{
	// 	AllowOrigins:     []string{"http://localhost:4000"},
	// 	AllowMethods:     []string{"GET", "POST", "OPTIONS"},
	// 	AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
	// 	ExposeHeaders:    []string{"Content-Length"},
	// 	AllowCredentials: true,
	// 	MaxAge:           12 * time.Hour,
	// }))
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, World!")
	})
	auth := r.Group("/api/auth")
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
	return r
}
