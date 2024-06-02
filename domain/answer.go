package domain

type Answer struct {
	Id         int
	QuestionId int
	AnswerId   int
	IsCorrect  bool
}
