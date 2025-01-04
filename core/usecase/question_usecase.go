package usecase

import (
	"github.com/google/uuid"
	"phishing-quest/adapter/repository"
	"phishing-quest/domain"
)

type QuestionUseCase struct {
	questionRepo repository.IQuestionRepository
}

func NewQuestionUseCase(questionRepo repository.IQuestionRepository) *QuestionUseCase {
	return &QuestionUseCase{questionRepo: questionRepo}
}

func (quc *QuestionUseCase) CreateQuestion(questionRequest *domain.Question) (*domain.Question, error) {
	question := &domain.Question{
		Id:            uuid.New(),
		CategoryId:    questionRequest.CategoryId,
		QuestionText:  questionRequest.QuestionText,
		CorrectAnswer: questionRequest.CorrectAnswer,
	}

	err := question.Validate()
	if err != nil {
		return nil, err
	}

	createdQuestion, err := quc.questionRepo.Create(question)
	if err != nil {
		return nil, err
	}

	return createdQuestion, nil
}

func (quc *QuestionUseCase) GetQuestion(id uuid.UUID) (*domain.Question, error) {
	question, err := quc.questionRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	return question, nil
}

func (quc *QuestionUseCase) UpdateQuestion(id uuid.UUID, questionRequest *domain.Question) (*domain.Question, error) {
	// Obter a pergunta existente
	question, err := quc.questionRepo.GetByID(id)
	if err != nil {
		return nil, err
	}

	if questionRequest.CategoryId != uuid.Nil {
		question.CategoryId = questionRequest.CategoryId
	}
	if questionRequest.QuestionText != "" {
		question.QuestionText = questionRequest.QuestionText
	}
	if questionRequest.CorrectAnswer != "" {
		question.CorrectAnswer = questionRequest.CorrectAnswer
	}

	err = question.Validate()
	if err != nil {
		return nil, err
	}

	updatedQuestion, err := quc.questionRepo.Update(question)
	if err != nil {
		return nil, err
	}

	return updatedQuestion, nil
}

func (quc *QuestionUseCase) DeleteQuestion(id uuid.UUID) error {
	_, err := quc.questionRepo.GetByID(id)
	if err != nil {
		return err
	}

	err = quc.questionRepo.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
