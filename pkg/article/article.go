package article

import (
	"assessment2/pkg/article/controller"
	"assessment2/pkg/article/usecase"
	"assessment2/pkg/article/repository"

	"gorm.io/gorm"
)

func InitHttpArticleController(db *gorm.DB) *controller.ArticleHTTPController {
	repo := repository.InitRepositoryArticle(db)
	uc := usecase.InitUsecaseArticle(repo)
	controller := controller.InitControllerArticle(uc)

	return controller
}
