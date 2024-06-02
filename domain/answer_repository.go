package domain

type AnswerRepository interface {
	CreateAnswer(answer *Answer) (*Answer, error)
	FindByQuestionId(questionId int) ([]*Answer, error)
}
