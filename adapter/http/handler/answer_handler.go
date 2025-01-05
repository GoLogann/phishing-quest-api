package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"phishing-quest/core/usecase"
	"phishing-quest/domain"
)

type AnswerHandler struct {
	answerUseCase *usecase.AnswerUseCase
}

func NewAnswerHandler(aqc *usecase.AnswerUseCase) *AnswerHandler {
	return &AnswerHandler{answerUseCase: aqc}
}

func (ah *AnswerHandler) CreateAnswer(c *gin.Context) {
	var answerDTO *domain.Answer
	if err := c.ShouldBindJSON(&answerDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdAnswer, err := ah.answerUseCase.CreateAnswer(answerDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdAnswer)
}

func (ah *AnswerHandler) GetAnswer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	answer, err := ah.answerUseCase.GetAnswer(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Answer not found"})
		return
	}

	c.JSON(http.StatusOK, answer)
}

func (ah *AnswerHandler) UpdateAnswer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var answerDTO *domain.Answer
	if err := c.ShouldBindJSON(&answerDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedAnswer, err := ah.answerUseCase.UpdateAnswer(id, answerDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedAnswer)
}

func (ah *AnswerHandler) DeleteAnswer(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = ah.answerUseCase.DeleteAnswer(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
