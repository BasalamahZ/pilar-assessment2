package controller

import (
	"assessment2/pkg/article/model"
	"assessment2/pkg/article/usecase"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type ArticleHTTPController struct {
	usecase usecase.UsecaseInterfaceArticle
}

func InitControllerArticle(uc usecase.UsecaseInterfaceArticle) *ArticleHTTPController {
	return &ArticleHTTPController{
		usecase: uc,
	}
}

func (uc *ArticleHTTPController) PostArticle(c *gin.Context) {
	var input model.Article
	err := c.ShouldBindJSON(&input)
	if err != nil {
		log.Println("error :", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err.Error(),
		})
		return
	}

	err_ := uc.usecase.NewArticle(input)
	if err_ != nil {
		log.Println(err_)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err_.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "article created",
		"data": gin.H{
			"article": input,
		},
	})
}

func (uc *ArticleHTTPController) DetailArticle(c *gin.Context) {
	articleId := c.Param("id")
	newArticleId, err := strconv.Atoi(articleId)
	if err != nil {
		return
	}
	article, err := uc.usecase.DetailArticle(newArticleId)
	if err != nil {
		log.Println("error :", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "success get article",
		"data": gin.H{
			"article": article,
		},
	})
}

func (uc *ArticleHTTPController) GetArticle(c *gin.Context) {
	var total_page float64
	limit_s, _ := c.GetQuery("limit")
	page_s, _ := c.GetQuery("page")

	// make pagination
	limit, err := strconv.Atoi(limit_s)
	if err != nil {
		limit = 5
	}

	page, err := strconv.Atoi(page_s)
	if err != nil {
		page = 2
	}

	if page > 1 {
		page = page*limit - limit
	} else {
		page = 0
	}

	articles, total, err := uc.usecase.GetAllArticle(limit, page)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"status":  http.StatusBadRequest,
			"message": "something wrong",
			"err":     err.Error(),
		})
		return
	}

	// make total page
	total_page = float64(total) / float64(limit)
	if total_page-float64(int(total_page)) > 0 {
		total_page = float64(int(total_page) + 1)
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "sucess get article",
		"data": gin.H{
			"article":    articles,
			"total_page": int(total_page),
			"total_data": total,
			"page":       page + 1,
			"limit":      limit,
		},
	})
}
