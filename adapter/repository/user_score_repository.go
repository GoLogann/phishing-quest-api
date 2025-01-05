package repository

import (
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"phishing-quest/domain"
)

type IUserScoreRepository interface {
	IRepository[domain.UserScore]
	IncrementScore(userID uuid.UUID, points int) error
	GetUserScore(userID uuid.UUID) (*domain.UserScore, error)
}

type UserScoreRepository struct {
	IRepository[domain.UserScore]
	db *gorm.DB
}

func NewUserScoreRepository(db *gorm.DB) IUserScoreRepository {
	return &UserScoreRepository{
		IRepository: NewRepository[domain.UserScore](db),
		db:          db,
	}
}

func (r *UserScoreRepository) IncrementScore(userID uuid.UUID, points int) error {
	logrus.Infof("Incrementando pontuação para o usuário com ID: %s, pontos: %d", userID, points)
	result := r.db.Model(&domain.UserScore{}).
		Where("user_id = ?", userID).
		Update("score", gorm.Expr("score + ?", points))
	if result.Error != nil {
		logrus.Errorf("Erro ao incrementar pontuação para o usuário com ID %s: %v", userID, result.Error)
		return result.Error
	}
	if result.RowsAffected == 0 {
		logrus.Warnf("Nenhum registro encontrado para incrementar pontuação para o usuário com ID: %s", userID)
		return gorm.ErrRecordNotFound
	}
	logrus.Infof("Pontuação incrementada com sucesso para o usuário com ID: %s", userID)
	return nil
}

func (r *UserScoreRepository) GetUserScore(userID uuid.UUID) (*domain.UserScore, error) {
	logrus.Infof("Buscando pontuação para o usuário com ID: %s", userID)
	var userScore domain.UserScore
	if err := r.db.Where("user_id = ?", userID).First(&userScore).Error; err != nil {
		logrus.Errorf("Erro ao buscar pontuação para o usuário com ID %s: %v", userID, err)
		return nil, err
	}
	logrus.Infof("Pontuação encontrada para o usuário com ID: %s", userID)
	return &userScore, nil
}
