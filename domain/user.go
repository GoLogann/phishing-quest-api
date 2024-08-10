package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

type User struct {
	Id           uuid.UUID `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username" validate:"required,min=1,max=255"`
	Email        string    `json:"email" validate:"required,email,max=255"`
	Password     string    `json:"password" validate:"required" gorm:"-"`
	PasswordHash string    `json:"-" validate:"required,min=1,max=255"`
	TotalScore   int       `json:"totalScore" validate:"gte=0"`
	CreatedAt    time.Time `json:"createdAt"`
	UpdatedAt    time.Time `json:"updatedAt"`
}

func (u *User) TableName() string {
	return "phishing_quest.users"
}
func (u *User) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}
