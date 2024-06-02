package service

import (
	"phishing-quest/internal/adapters/database"
	"phishing-quest/internal/domain/entity"
	"phishing-quest/internal/domain/repository"
)

type QuestionService struct {
	repo repository.QuestionRepository
}

func NewQuestionService(repo *database.QuestionRepository) *QuestionService {
	return &QuestionService{repo: repo}
}

func (s *QuestionService) CreateQuestion(categoryId int, questionText, correctAnswer string) (*entity.Question, error) {
	question := &entity.Question{
		CategoryId:    categoryId,
		QuestionText:  questionText,
		CorrectAnswer: correctAnswer,
	}
	_, err := s.repo.CreateQuestion(question)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (s *QuestionService) GetQuestionByID(id int) (*entity.Question, error) {
	return s.repo.FindQuestionById(id)
}

//func (s *QuestionService) GetQuestionsByCategory(categoryID int) ([]*entity.Question, error) {
//    return s.repo.FindQuestionByCategoryId(categoryID)
//}
