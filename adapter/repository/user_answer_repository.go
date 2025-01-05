package repository

import (
	"gorm.io/gorm"
	"phishing-quest/domain"
)

type IUserAnswerRepository interface {
	IRepository[domain.UserAnswer]
}

type UserAnswerRepository struct {
	IRepository[domain.UserAnswer]
	db *gorm.DB
}

func NewUserAnswerRepository(db *gorm.DB) IUserAnswerRepository {
	return &UserAnswerRepository{
		IRepository: NewRepository[domain.UserAnswer](db),
		db:          db,
	}
}
