package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"phishing-quest/core/usecase"
	"strconv"
)

type RankingHandler struct {
	rankingUseCase *usecase.RankingUseCase
}

func NewRankingHandler(ruc *usecase.RankingUseCase) *RankingHandler {
	return &RankingHandler{rankingUseCase: ruc}
}

func (rh *RankingHandler) GetGlobalRanking(c *gin.Context) {
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))
	offset, _ := strconv.Atoi(c.DefaultQuery("offset", "0"))

	ranking, err := rh.rankingUseCase.GetGlobalRanking(limit, offset)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch ranking"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"ranking": ranking})
}
