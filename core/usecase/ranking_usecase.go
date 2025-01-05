package usecase

import (
	"phishing-quest/adapter/repository"
	"phishing-quest/domain"
)

type RankingUseCase struct {
	rankingRepo repository.IRankingRepository
}

// NewRankingUseCase cria uma nova instância de RankingUseCase
func NewRankingUseCase(rankingRepo repository.IRankingRepository) *RankingUseCase {
	return &RankingUseCase{rankingRepo: rankingRepo}
}

// GetGlobalRanking retorna o ranking global, paginado
func (ruc *RankingUseCase) GetGlobalRanking(limit, offset int) ([]domain.UserScore, error) {
	ranking, err := ruc.rankingRepo.GetGlobalRanking(limit, offset)
	if err != nil {
		return nil, err
	}

	return ranking, nil
}
