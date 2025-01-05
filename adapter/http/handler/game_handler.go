package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phishing-quest/core/usecase"
	"phishing-quest/dto"
)

type GameHandler struct {
	gameUseCase *usecase.GameUseCase
}

func NewGameHandler(guc *usecase.GameUseCase) *GameHandler {
	return &GameHandler{gameUseCase: guc}
}

func (gh *GameHandler) SubmitAnswer(c *gin.Context) {
	var answerDTO dto.SubmitAnswerDTO
	if err := c.ShouldBindJSON(&answerDTO); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := gh.gameUseCase.ProcessAnswer(&answerDTO)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}
