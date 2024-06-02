package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phishing-quest/internal/domain/service"
	"strconv"
)

type QuestionHandler struct {
	questionService *service.QuestionService
}

func NewQuestionHandler(questionService *service.QuestionService) *QuestionHandler {
	return &QuestionHandler{questionService: questionService}
}

func (h *QuestionHandler) CreateQuestion(c *gin.Context) {
	var question struct {
		CategoryID    int    `json:"category_id"`
		QuestionText  string `json:"question_text"`
		CorrectAnswer string `json:"correct_answer"`
	}
	if err := c.BindJSON(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	createdQuestion, err := h.questionService.CreateQuestion(question.CategoryID, question.QuestionText, question.CorrectAnswer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, createdQuestion)
}

func (h *QuestionHandler) GetQuestionByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid question ID"})
		return
	}

	question, err := h.questionService.GetQuestionByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, question)
}

//func (h *QuestionHandler) GetQuestionsByCategory(c *gin.Context) {
//	categoryID, err := strconv.Atoi(c.Param("category_id"))
//	if err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid category ID"})
//		return
//	}
//
//	questions, err := h.questionService.GetQuestionsByCategory(categoryID)
//	if err != nil {
//		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
//		return
//	}
//
//	c.JSON(http.StatusOK, questions)
//}
