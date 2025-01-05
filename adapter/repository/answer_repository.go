package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"phishing-quest/domain"
)

type IAnswerRepository interface {
	IRepository[domain.Answer]
	GetByQuestionID(questionID uuid.UUID) ([]*domain.Answer, error)
}

type AnswerRepository struct {
	IRepository[domain.Answer]
	db *gorm.DB
}

func NewAnswerRepository(db *gorm.DB) IAnswerRepository {
	return &AnswerRepository{
		IRepository: NewRepository[domain.Answer](db),
		db:          db,
	}
}

func (ar *AnswerRepository) GetByQuestionID(questionID uuid.UUID) ([]*domain.Answer, error) {
	var answers []*domain.Answer
	if err := ar.db.Where("question_id = ?", questionID).Find(&answers).Error; err != nil {
		return nil, err
	}
	return answers, nil
}
