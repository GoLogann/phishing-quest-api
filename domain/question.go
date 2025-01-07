package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"phishing-quest/dto"
)

type Question struct {
	Id            uuid.UUID `json:"id" gorm:"primaryKey"`
	CategoryId    uuid.UUID `json:"categoryId" validate:"required"`
	QuestionText  string    `json:"questionText" validate:"required"`
	CorrectAnswer string    `json:"correctAnswer" validate:"required"`
}

func (q *Question) TableName() string {
	return "phishing_quest.questions"
}

func (q *Question) Validate() error {
	validate := validator.New()
	return validate.Struct(q)
}

func (q *Question) ToDTO() *dto.QuestionDTO {
	return &dto.QuestionDTO{
		QuestionId:   q.Id,
		QuestionText: q.QuestionText,
	}
}
