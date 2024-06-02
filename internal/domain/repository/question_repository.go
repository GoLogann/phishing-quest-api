package repository

import "phishing-quest/internal/domain/entity"

type QuestionRepository interface {
	createQuestion(question *entity.Question) (*entity.Question, error)
	FindQuestionById(id int) (*entity.Question, error)
	FindQuestionByCategoryId(categoryId int) (*entity.Question, error)
}
