package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"phishing-quest/core/usecase"
	"phishing-quest/domain"
	"phishing-quest/dto"
)

type QuestionHandler struct {
	questionUseCase *usecase.QuestionUseCase
}

func NewQuestionHandler(quc *usecase.QuestionUseCase) *QuestionHandler {
	return &QuestionHandler{questionUseCase: quc}
}

func (qh *QuestionHandler) CreateQuestion(c *gin.Context) {
	var questionDTO *domain.Question
	if err := c.ShouldBindJSON(&questionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdQuestion, err := qh.questionUseCase.CreateQuestion(questionDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, createdQuestion)
}

func (qh *QuestionHandler) ListAnswersByQuestion(c *gin.Context) {
	idParam := c.Param("id")
	questionID, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Question ID"})
		return
	}

	answers, err := qh.questionUseCase.GetAnswersByQuestionID(questionID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Answers not found"})
		return
	}

	var answerDTOs []*dto.AnswerDTO
	for _, answer := range answers {
		answerDTOs = append(answerDTOs, answer.ToDTO())
	}

	response := dto.QuestionAnswersDTO{
		QuestionId: questionID,
		Answers:    answerDTOs,
	}

	c.JSON(http.StatusOK, response)
}

func (qh *QuestionHandler) GetQuestion(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	question, err := qh.questionUseCase.GetQuestion(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Question not found"})
		return
	}

	c.JSON(http.StatusOK, question)
}

func (qh *QuestionHandler) UpdateQuestion(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	var questionDTO *domain.Question
	if err := c.ShouldBindJSON(&questionDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updatedQuestion, err := qh.questionUseCase.UpdateQuestion(id, questionDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, updatedQuestion)
}

func (qh *QuestionHandler) DeleteQuestion(c *gin.Context) {
	idParam := c.Param("id")
	id, err := uuid.Parse(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return
	}

	err = qh.questionUseCase.DeleteQuestion(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusNoContent, nil)
}
