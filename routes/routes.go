package routes

import (
	"assessment2/pkg/article"
	"assessment2/pkg/user"
	"assessment2/utils/helper/auth"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func InitHttpRoute(g *gin.Engine, db *gorm.DB) {
	apiGroup := g.Group("/api")
	pingGroup := g.Group("ping")

	pingGroup.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	userController := user.InitHttpUserController(db)
	articleController := article.InitHttpArticleController(db)
	apiGroup.POST("/register", userController.Register)
	apiGroup.POST("/api/token", userController.Login)
	apiGroup.POST("/api/token/refresh", userController.RefreshAccessToken)
	apiGroup.Use(auth.MiddlewareLogin())
	{
		apiGroup.GET("/profile", userController.ViewProfile)
		apiGroup.GET("/article", articleController.GetArticle)
		apiGroup.GET("/article/:id", articleController.DetailArticle)
		apiGroup.POST("/article", articleController.PostArticle)
	}
}
