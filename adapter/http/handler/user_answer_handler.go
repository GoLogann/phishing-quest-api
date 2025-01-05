package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"phishing-quest/core/usecase"
	"phishing-quest/domain"
)

type UserAnswerHandler struct {
	userAnswerUseCase *usecase.UserAnswerUseCase
}

func NewUserAnswerHandler(uauc *usecase.UserAnswerUseCase) *UserAnswerHandler {
	return &UserAnswerHandler{userAnswerUseCase: uauc}
}

func (uah *UserAnswerHandler) CreateUserAnswer(c *gin.Context) {
	var userAnswer domain.UserAnswer
	if err := c.ShouldBindJSON(&userAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdUserAnswer, err := uah.userAnswerUseCase.CreateUserAnswer(&userAnswer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdUserAnswer)
}

func (uah *UserAnswerHandler) GetUserAnswer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	userAnswer, err := uah.userAnswerUseCase.GetUserAnswerByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "UserAnswer not found"})
		return
	}

	c.JSON(http.StatusOK, userAnswer)
}

func (uah *UserAnswerHandler) ListUserAnswers(c *gin.Context) {
	userAnswers, err := uah.userAnswerUseCase.ListUserAnswers()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, userAnswers)
}

func (uah *UserAnswerHandler) UpdateUserAnswer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var userAnswer domain.UserAnswer
	if err := c.ShouldBindJSON(&userAnswer); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedUserAnswer, err := uah.userAnswerUseCase.UpdateUserAnswer(id, &userAnswer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedUserAnswer)
}

func (uah *UserAnswerHandler) DeleteUserAnswer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = uah.userAnswerUseCase.DeleteUserAnswer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
