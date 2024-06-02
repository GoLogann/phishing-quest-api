package repository

import "phishing-quest/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindUserById(id string) (*entity.User, error)
	DeleteUser(id string) error
	UpdateUser(user *entity.User) error
}
