package usecase

import (
	"github.com/google/uuid"
	"phishing-quest/adapter/repository"
	"phishing-quest/domain"
)

type UserAnswerUseCase struct {
	userAnswerRepo repository.IUserAnswerRepository
}

func NewUserAnswerUseCase(repo repository.IUserAnswerRepository) *UserAnswerUseCase {
	return &UserAnswerUseCase{userAnswerRepo: repo}
}

func (uauc *UserAnswerUseCase) CreateUserAnswer(userAnswer *domain.UserAnswer) (*domain.UserAnswer, error) {
	userAnswer.UserAnswerId = uuid.New()
	return uauc.userAnswerRepo.Create(userAnswer)
}

func (uauc *UserAnswerUseCase) GetUserAnswerByID(id uuid.UUID) (*domain.UserAnswer, error) {
	return uauc.userAnswerRepo.GetByID(id)
}

func (uauc *UserAnswerUseCase) ListUserAnswers() ([]*domain.UserAnswer, error) {
	return uauc.userAnswerRepo.GetAll()
}

func (uauc *UserAnswerUseCase) UpdateUserAnswer(id uuid.UUID, userAnswer *domain.UserAnswer) (*domain.UserAnswer, error) {
	existing, err := uauc.userAnswerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	existing.AnswerId = userAnswer.AnswerId
	existing.IsCorrect = userAnswer.IsCorrect
	existing.Timestamp = userAnswer.Timestamp
	existing.AiRating = userAnswer.AiRating

	return uauc.userAnswerRepo.Update(existing)
}

func (uauc *UserAnswerUseCase) DeleteUserAnswer(id uuid.UUID) error {
	return uauc.userAnswerRepo.Delete(id)
}
