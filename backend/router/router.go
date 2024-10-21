package router

import (
	"backend/controllers"
	"backend/middlewares"

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
	//
	auth.POST("/register", controllers.Register)
	auth.POST("/login", controllers.Login)
	//

	api := r.Group("/api")
	api.GET("/exchangeRates", controllers.GetExchangeRate)
	api.Use(middlewares.AuthMiddleware())
	//
	api.POST("/exchangeRates", controllers.CreateExchangeRate)
	api.POST("/articles", controllers.CreateArticle)
	api.GET("/articles", controllers.GetArticles)
	api.GET("/articles/:id", controllers.GetArticleByID)
	//
	api.POST("/articles/:id/like", controllers.LikeArticle)
	api.GET("/articles/:id/likes", controllers.GetArticleLikes)
	return r
}
