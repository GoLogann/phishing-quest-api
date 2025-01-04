package repository

import (
	"gorm.io/gorm"
	"phishing-quest/domain"
)

type IUserRepository interface {
	IRepository[domain.User]
	GetByEmail(email string) (*domain.User, error)
}

type UserRepository struct {
	IRepository[domain.User]
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		IRepository: NewRepository[domain.User](db),
		db:          db,
	}
}

func (r *UserRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
