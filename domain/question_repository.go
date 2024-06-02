package domain

type QuestionRepository interface {
	createQuestion(question *Question) (*Question, error)
	FindQuestionById(id int) (*Question, error)
	FindQuestionByCategoryId(categoryId int) (*Question, error)
}
