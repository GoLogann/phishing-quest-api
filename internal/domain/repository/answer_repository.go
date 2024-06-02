package repository

import "phishing-quest/internal/domain/entity"

type AnswerRepository interface {
	CreateAnswer(answer *entity.Answer) (*entity.Answer, error)
	FindByQuestionId(questionId int) ([]*entity.Answer, error)
}
