package repository

import "phishing-quest/internal/domain/entity"

type UserRepository interface {
	Create(user *entity.User) error
	FindUserById(id int) (*entity.User, error)
	DeleteUser(id int) error
	UpdateUser(user *entity.User) error
}
