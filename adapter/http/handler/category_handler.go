package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phishing-quest/core/usecase"
	"phishing-quest/domain"
)

type CategoryHandler struct {
	categoryUseCase *usecase.CategoryUseCase
}

func NewCategoryHandler(cuc *usecase.CategoryUseCase) *CategoryHandler {
	return &CategoryHandler{categoryUseCase: cuc}
}

func (ch *CategoryHandler) CreateCategory(c *gin.Context) {
	var categoryDTO *domain.Category
	if err := c.ShouldBindJSON(&categoryDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdCategory, err := ch.categoryUseCase.CreateCategory(categoryDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdCategory)
}

func (ch *CategoryHandler) ListCategory(c *gin.Context) {
	categories, err := ch.categoryUseCase.ListCategories()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, categories)
}
