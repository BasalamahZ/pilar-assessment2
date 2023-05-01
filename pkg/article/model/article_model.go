package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	// ArticleId        int            `json:"article_id" gorm:"primary_key;auto_increment"`
	ArticleTitle     string         `json:"article_title"`
	ArtilceContent   string         `json:"article_content"`
}