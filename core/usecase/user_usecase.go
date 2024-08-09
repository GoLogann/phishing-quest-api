package usecase

import (
	"errors"
	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"phishing-quest/adapter/repository"
	"phishing-quest/domain"
	"phishing-quest/dto"
	"time"
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

	hashedPassword, err := uc.HashPassword(userRequest.Password)
	if err != nil {
		return nil, errors.New("erro ao gerar hash da senha")
	}

	user := &domain.User{
		Id:           uuid.New(),
		Username:     userRequest.Username,
		Email:        userRequest.Email,
		Password:     userRequest.Password,
		PasswordHash: hashedPassword,
		TotalScore:   0,
		CreatedAt:    time.Now(),
	}

	err = user.Validate()
	if err != nil {
		return nil, err
	}

	if err = uc.userRepo.Create(user); err != nil {
		return nil, err
	}

	return user, nil
}

func (uc *UserUseCase) Login(userRequest *dto.UserLoginDTO) (*domain.User, error) {
	user, err := uc.userRepo.GetByEmail(userRequest.Email)
	if err != nil {
		return nil, errors.New("usuário não encontrado")
	}

	if !uc.CheckPasswordHash(userRequest.Password, user.PasswordHash) {
		return nil, errors.New("senha incorreta")
	}

	return user, nil
}

// UpdatePassword atualiza a senha do usuário
func (uc *UserUseCase) UpdatePassword(user *domain.User, newPasswordHash string) {
	user.PasswordHash = newPasswordHash
	user.UpdatedAt = time.Now()
}

// AddScore adiciona uma quantidade específica de pontos ao total do usuário
func (uc *UserUseCase) AddScore(user *domain.User, score int) {
	user.TotalScore += score
	user.UpdatedAt = time.Now()
}

// HashPassword recebe uma senha e retorna o hash dela.
func (uc *UserUseCase) HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}

func (uc *UserUseCase) CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
