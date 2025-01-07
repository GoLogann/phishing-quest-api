package repository

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
	"phishing-quest/domain"
)

type IQuestionRepository interface {
	IRepository[domain.Question]
	GetByCategoryID(questionID uuid.UUID) ([]*domain.Question, error)
}

type QuestionRepository struct {
	IRepository[domain.Question]
	db *gorm.DB
}

func NewQuestionRepository(db *gorm.DB) IQuestionRepository {
	return &QuestionRepository{
		IRepository: NewRepository[domain.Question](db),
		db:          db,
	}
}

func (qr *QuestionRepository) GetByCategoryID(categoryID uuid.UUID) ([]*domain.Question, error) {
	var questions []*domain.Question
	if err := qr.db.Where("category_id = ?", categoryID).Find(&questions).Error; err != nil {
		return nil, err
	}
	return questions, nil
}
