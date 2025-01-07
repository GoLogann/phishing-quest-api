package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"phishing-quest/core/usecase"
	"phishing-quest/domain"
	"phishing-quest/dto"
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

func (ch *CategoryHandler) ListQuestionsByCategory(c *gin.Context) {
	idParam := c.Param("category_id")
	categoryID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Question ID"})
		return
	}

	questions, err := ch.categoryUseCase.GetQuestionsByCategoryID(categoryID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Answers not found"})
		return
	}

	var questionDTOs []*dto.QuestionDTO
	for _, question := range questions {
		questionDTOs = append(questionDTOs, question.ToDTO())
	}

	response := dto.CategoryQuestionsDTO{
		CategoryId: categoryID,
		Questions:  questionDTOs,
	}

	c.JSON(http.StatusOK, response)
}
