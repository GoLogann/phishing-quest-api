package dto

import "github.com/google/uuid"

type UserLoginDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginResponseDTO struct {
	Id         uuid.UUID `json:"id"`
	Username   string    `json:"username"`
	Email      string    `json:"email"`
	TotalScore int       `json:"totalScore"`
}
