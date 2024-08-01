package usecase

import (
	"errors"
	"phishing-quest/adapter/repository"
	"phishing-quest/domain"
)

// UserUseCase representa os casos de uso relacionados a usuários
type UserUseCase struct {
	userRepo repository.UserRepository
}

// NewUserUseCase cria um novo caso de uso para usuários
func NewUserUseCase(userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{userRepo: userRepo}
}

// CreateUser cria um novo usuário
func (uc *UserUseCase) CreateUser(userRequest *domain.User) (*domain.User, error) {
	existingUser, _ := uc.userRepo.GetByEmail(userRequest.Email)
	if existingUser != nil {
		return nil, errors.New("email já está em uso")
	}

	user := &domain.User{
		Username:     userRequest.Username,
		Email:        userRequest.Email,
		PasswordHash: userRequest.PasswordHash,
		XP:           0,
		TotalScore:   0,
		//CreatedAt:    time.Now(),
		//UpdatedAt:    time.Now(),
	}

	if err := uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

// UpdatePassword atualiza a senha do usuário
func (uc *UserUseCase) UpdatePassword(user *domain.User, newPasswordHash string) {
	user.PasswordHash = newPasswordHash
	//user.UpdatedAt = time.Now()
	// lógica adicional, como persistência, pode ser adicionada aqui
}

// AddXP adiciona uma quantidade específica de XP ao usuário
func (uc *UserUseCase) AddXP(user *domain.User, amount int) {
	user.XP += amount
	//user.UpdatedAt = time.Now()
	// lógica adicional, como persistência, pode ser adicionada aqui
}

// AddScore adiciona uma quantidade específica de pontos ao total do usuário
func (uc *UserUseCase) AddScore(user *domain.User, score int) {
	user.TotalScore += score
	//user.UpdatedAt = time.Now()
	// lógica adicional, como persistência, pode ser adicionada aqui
}
