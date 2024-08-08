package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"time"
)

// User representa a entidade de um usuário no sistema
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

// NewUser é um construtor para criar um novo usuário
func NewUser(username, email, passwordHash string) *User {
	return &User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
		TotalScore:   0,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
}
