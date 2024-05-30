package handler

import (
	"strconv"
	"testapp1/pkg/domain"

	"github.com/gin-gonic/gin"
)

type articleHandler struct {
	usecase domain.ArticleUsecase
}

func (h *articleHandler) GetAll(c *gin.Context) {
	articles, err := h.usecase.GetAll()
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, articles)
}

func NewArticleHandler(usecase domain.ArticleUsecase) *articleHandler {
	return &articleHandler{
		usecase: usecase,
	}
}

func (h *articleHandler) GetByID(c *gin.Context) {
	id := c.Param("id")
	// Convert id from string to int
	idInt, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(400, gin.H{
			"error": "Invalid ID",
		})
		return
	}

	article, err := h.usecase.GetByID(idInt)
	if err != nil {
		c.JSON(500, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(200, article)
}
