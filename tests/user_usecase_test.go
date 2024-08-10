package tests

import (
	"testing"
	"time"

	"phishing-quest/core/usecase"
	"phishing-quest/domain"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) Create(user *domain.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserRepository) GetByID(id int) (*domain.User, error) {
	args := m.Called(id)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockUserRepository) GetByEmail(email string) (*domain.User, error) {
	args := m.Called(email)
	if args.Get(0) != nil {
		return args.Get(0).(*domain.User), args.Error(1)
	}
	return nil, args.Error(1)
}

func TestUserUseCase_CreateUser(t *testing.T) {
	t.Run("user creation success", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		uc := usecase.NewUserUseCase(mockRepo)

		userRequest := &domain.User{
			Username: "testuser",
			Email:    "test@example.com",
			Password: "password123",
		}

		mockRepo.On("GetByEmail", userRequest.Email).Return(nil, gorm.ErrRecordNotFound)
		mockRepo.On("Create", mock.AnythingOfType("*domain.User")).Return(nil)

		user, err := uc.CreateUser(userRequest)

		assert.NoError(t, err)
		assert.NotNil(t, user)
		assert.Equal(t, userRequest.Username, user.Username)
		assert.Equal(t, userRequest.Email, user.Email)
		mockRepo.AssertExpectations(t)
	})

	t.Run("user creation fails when email already exists", func(t *testing.T) {
		mockRepo := new(MockUserRepository)
		uc := usecase.NewUserUseCase(mockRepo)

		existingUser := &domain.User{
			Id:        uuid.New(),
			Username:  "existinguser",
			Email:     "test@example.com",
			CreatedAt: time.Now(),
		}

		mockRepo.On("GetByEmail", existingUser.Email).Return(existingUser, nil)

		userRequest := &domain.User{
			Username: "newuser",
			Email:    "test@example.com",
			Password: "password123",
		}

		user, err := uc.CreateUser(userRequest)

		assert.Nil(t, user)
		assert.EqualError(t, err, "email já está em uso")
		mockRepo.AssertExpectations(t)
	})
}
