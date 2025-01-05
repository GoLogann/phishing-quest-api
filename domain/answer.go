package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"phishing-quest/dto"
)

type Answer struct {
	Id         uuid.UUID `json:"id" gorm:"primaryKey"`
	QuestionId uuid.UUID `json:"questionId" validate:"required"`
	AnswerText string    `json:"answerText" validate:"required"`
	IsCorrect  bool      `json:"isCorrect" validate:"required"`
}

func (a *Answer) TableName() string {
	return "phishing_quest.answers"
}

func (a *Answer) Validate() error {
	validate := validator.New()
	return validate.Struct(a)
}

func (a *Answer) ToDTO() *dto.AnswerDTO {
	return &dto.AnswerDTO{
		Id:         a.Id,
		AnswerText: a.AnswerText,
		IsCorrect:  a.IsCorrect,
	}
}
