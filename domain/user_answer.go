package domain

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type UserAnswer struct {
	UserAnswerId uuid.UUID `json:"userAnswerId" gorm:"primaryKey"`
	UserId       uuid.UUID `json:"userId" validate:"required"`
	QuestionId   uuid.UUID `json:"questionId" validate:"required"`
	AnswerId     uuid.UUID `json:"answerId" validate:"required"`
	IsCorrect    *bool     `json:"isCorrect"`
	Timestamp    time.Time `json:"timestamp" gorm:"default:current_timestamp"`
	AiRating     *int      `json:"aiRating" validate:"omitempty,gte=1,lte=5"`
}

func (ua *UserAnswer) TableName() string {
	return "phishing_quest.user_answers"
}

func (ua *UserAnswer) Validate() error {
	validate := validator.New()
	return validate.Struct(ua)
}
