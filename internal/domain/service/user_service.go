package service

import (
	"phishing-quest/internal/domain/entity"
	"phishing-quest/internal/domain/repository"
)

type UserService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(username, email, passwordHash string) (*entity.User, error) {
	user := &entity.User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}
	err := s.repo.Create(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByID(id string) (*entity.User, error) {
	return s.repo.FindUserById(id)
}
