package repository

import (
	"gorm.io/gorm"
	"phishing-quest/domain"
)

// IRankingRepository define os métodos do repositório de ranking
type IRankingRepository interface {
	GetGlobalRanking(limit int, offset int) ([]domain.UserScore, error)
}

// RankingRepository implementa IRankingRepository
type RankingRepository struct {
	db *gorm.DB
}

// NewRankingRepository cria uma nova instância de RankingRepository
func NewRankingRepository(db *gorm.DB) IRankingRepository {
	return &RankingRepository{
		db: db,
	}
}

// GetGlobalRanking retorna o ranking global ordenado pela pontuação total
func (rr *RankingRepository) GetGlobalRanking(limit int, offset int) ([]domain.UserScore, error) {
	var scores []domain.UserScore
	err := rr.db.
		Model(&domain.UserScore{}).
		Select("user_id, SUM(score) AS total_score").
		Group("user_id").
		Order("total_score DESC").
		Limit(limit).
		Offset(offset).
		Find(&scores).Error

	if err != nil {
		return nil, err
	}
	return scores, nil
}
