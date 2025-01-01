package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type UserScore struct {
	Id        uuid.UUID `json:"id" gorm:"primaryKey"`
	UserId    uuid.UUID `json:"userId" validate:"required"`
	Score     int       `json:"score" validate:"required,gte=0"`
	Timestamp time.Time `json:"timestamp" gorm:"default:current_timestamp"`
}

func (us *UserScore) TableName() string {
	return "phishing_quest.user_scores"
}

func (us *UserScore) Validate() error {
	validate := validator.New()
	return validate.Struct(us)
}

//CREATE VIEW phishing_quest.global_ranking AS
//SELECT user_id, SUM(score) AS total_score
//FROM phishing_quest.user_scores
//GROUP BY user_id
//ORDER BY total_score DESC;
