package dto

import "github.com/google/uuid"

type QuestionDTO struct {
	QuestionId   uuid.UUID `json:"questionId"`
	QuestionText string    `json:"questionText"`
}

type CategoryQuestionsDTO struct {
	CategoryId uuid.UUID      `json:"categoryId"`
	Questions  []*QuestionDTO `json:"questions"`
}
