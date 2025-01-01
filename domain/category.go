package domain

import (
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

type Category struct {
	Id           uuid.UUID `json:"id" gorm:"primaryKey"`
	CategoryName string    `json:"categoryName" validate:"required,min=1,max=255"`
}

func (c *Category) TableName() string {
	return "phishing_quest.categories"
}

func (c *Category) Validate() error {
	validate := validator.New()
	return validate.Struct(c)
}
