package repository

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"phishing-quest/domain"
)

type UserRepository interface {
	Create(user *domain.User) error
	GetByID(id int) (*domain.User, error)
	GetByEmail(email string) (*domain.User, error)
}

type userRepository struct {
	db *gorm.DB
}

// NewUserRepository cria uma nova instância de userRepository
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *domain.User) error {
	logrus.WithFields(logrus.Fields{
		"username": user.Username,
		"email":    user.Email,
	}).Info("Create new user")

	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id int) (*domain.User, error) {
	var user domain.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetByEmail(email string) (*domain.User, error) {
	var user domain.User
	if err := r.db.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
