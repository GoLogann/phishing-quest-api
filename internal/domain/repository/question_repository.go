package repository

import "phishing-quest/internal/domain/entity"

type QuestionRepository interface {
	CreateQuestion(question *entity.Question) (*entity.Question, error)
	FindQuestionById(id int) (*entity.Question, error)
	FindQuestionByCategoryId(categoryId int) (*entity.Question, error)
}
