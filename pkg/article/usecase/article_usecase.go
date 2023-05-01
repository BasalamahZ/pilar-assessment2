package usecase

import (
	"assessment2/pkg/article/model"
	"assessment2/pkg/article/repository"
)

type UsecaseInterfaceArticle interface {
	NewArticle(input model.Article) error
	DetailArticle(articleId int) (model.Article, error)
	GetAllArticle(limit, offset int) ([]model.Article, int, error)
}

type usecaseArticle struct {
	repository repository.RepositoryInterfaceArticle
}

func InitUsecaseArticle(repository repository.RepositoryInterfaceArticle) UsecaseInterfaceArticle {
	return &usecaseArticle{
		repository: repository,
	}
}

// NewArticle implements UsecaseInterfaceArticle
func (u *usecaseArticle) NewArticle(input model.Article) error {
	err := u.repository.CreateNewArticle(input)
	return err
}

// DetailArticle implements UsecaseInterfaceArticle
func (u *usecaseArticle) DetailArticle(articleId int) (model.Article, error) {
	article, err := u.repository.DetailArticle(articleId)
	return article, err
}

// GetAllArticle implements UsecaseInterfaceArticle
func (u *usecaseArticle) GetAllArticle(limit int, offset int) ([]model.Article, int, error) {
	articles, _ := u.repository.GetAllArticle(limit, offset)
	total, err := u.repository.GetTotalArticle()
	return articles, total, err
}