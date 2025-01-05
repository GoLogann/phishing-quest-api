package repository

import (
	"gorm.io/gorm"
	"phishing-quest/domain"
)

type IQuestionRepository interface {
	IRepository[domain.Question]
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
