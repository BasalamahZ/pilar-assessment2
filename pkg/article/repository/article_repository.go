package repository

import (
	"assessment2/pkg/article/model"

	"gorm.io/gorm"
)

type RepositoryInterfaceArticle interface {
	CreateNewArticle(input model.Article) error
	DetailArticle(articleId int) (model.Article, error)
	GetAllArticle(limit, offset int) ([]model.Article, error)
	GetTotalArticle() (int, error)
}

type repositoryArticle struct {
	db *gorm.DB
}

func InitRepositoryArticle(db *gorm.DB) RepositoryInterfaceArticle {
	db.AutoMigrate(&model.Article{})
	return &repositoryArticle{
		db: db,
	}
}

// CreateNewArticle implements RepositoryInterfaceArticle
func (r *repositoryArticle) CreateNewArticle(input model.Article) error {
	if err := r.db.Table("articles").Create(&input).Error; err != nil {
		return err
	}

	return nil
}

// DetailArticle implements RepositoryInterfaceArticle
func (r *repositoryArticle) DetailArticle(articleId int) (model.Article, error) {
	var article model.Article
	get := r.db.Table("articles").Where("ID = ?", articleId).Find(&article)

	if err := get.Error; err != nil {
		return article, err
	}

	return article, nil
}

// GetAllArticle implements RepositoryInterfaceArticle
func (r *repositoryArticle) GetAllArticle(limit int, offset int) ([]model.Article, error) {
	var articles []model.Article
	get := r.db.Table("articles").Limit(limit).Offset(offset).Find(&articles)
	if err := get.Error; err != nil {
		return articles, err
	}

	return articles, nil
}

// GetTotalArticle implements RepositoryInterfaceArticle
func (r *repositoryArticle) GetTotalArticle() (int, error) {
	var total int64
	get := r.db.Table("articles").Count(&total)

	if err := get.Error; err != nil {
		return int(total), err
	}

	return int(total), nil
}
