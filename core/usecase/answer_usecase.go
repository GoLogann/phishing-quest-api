package usecase

import (
	"github.com/google/uuid"
	"phishing-quest/adapter/repository"
	"phishing-quest/domain"
)

type AnswerUseCase struct {
	answerRepo repository.IAnswerRepository
}

func NewAnswerUseCase(answerRepo repository.IAnswerRepository) *AnswerUseCase {
	return &AnswerUseCase{answerRepo: answerRepo}
}

func (auc *AnswerUseCase) CreateAnswer(answerRequest *domain.Answer) (*domain.Answer, error) {
	answer := &domain.Answer{
		Id:         uuid.New(),
		QuestionId: answerRequest.QuestionId,
		AnswerText: answerRequest.AnswerText,
		IsCorrect:  answerRequest.IsCorrect,
	}

	err := answer.Validate()
	if err != nil {
		return nil, err
	}

	createdAnswer, err := auc.answerRepo.Create(answer)
	if err != nil {
		return nil, err
	}

	return createdAnswer, nil
}

func (auc *AnswerUseCase) GetAnswer(id uuid.UUID) (*domain.Answer, error) {
	answer, err := auc.answerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return answer, nil
}

func (auc *AnswerUseCase) UpdateAnswer(id uuid.UUID, answerRequest *domain.Answer) (*domain.Answer, error) {
	answer, err := auc.answerRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if answerRequest.AnswerText != "" {
		answer.AnswerText = answerRequest.AnswerText
	}
	answer.IsCorrect = answerRequest.IsCorrect

	err = answer.Validate()
	if err != nil {
		return nil, err
	}

	updatedAnswer, err := auc.answerRepo.Update(answer)
	if err != nil {
		return nil, err
	}

	return updatedAnswer, nil
}

func (auc *AnswerUseCase) DeleteAnswer(id uuid.UUID) error {
	_, err := auc.answerRepo.GetByID(id)
	if err != nil {
		return err
	}

	err = auc.answerRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
