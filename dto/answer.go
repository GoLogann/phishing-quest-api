package dto

import "github.com/google/uuid"

type QuestionAnswersDTO struct {
	QuestionId uuid.UUID    `json:"questionId"`
	Answers    []*AnswerDTO `json:"answers"`
}

type AnswerDTO struct {
	Id         uuid.UUID `json:"id"`
	AnswerText string    `json:"answerText"`
	IsCorrect  bool      `json:"isCorrect"`
}
