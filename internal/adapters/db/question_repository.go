package db

import "database/sql"

type QuestionRepository struct {
	db *sql.DB
}

func NewQuestionRepository(db *sql.DB) *QuestionRepository {
	return &QuestionRepository{db: db}
}

func (r *QuestionRepository) Save(question *domain.Question) error {
	_, err := r.db.Exec("INSERT INTO questions (category_id, question_text, correct_answer) VALUES ($1, $2, $3)", question.CategoryID, question.QuestionText, question.CorrectAnswer)
	return err
}

func (r *QuestionRepository) FindByID(id int) (*domain.Question, error) {
	row := r.db.QueryRow("SELECT question_id, category_id, question_text, correct_answer FROM questions WHERE question_id = $1", id)
	question := &domain.Question{}
	err := row.Scan(&question.ID, &question.CategoryID, &question.QuestionText, &question.CorrectAnswer)
	if err != nil {
		return nil, err
	}
	return question, nil
}

func (r *QuestionRepository) FindByCategoryID(categoryID int) ([]*domain.Question, error) {
	rows, err := r.db.Query("SELECT question_id, category_id, question_text, correct_answer FROM questions WHERE category_id = $1", categoryID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var questions []*domain.Question
	for rows.Next() {
		question := &domain.Question{}
		if err := rows.Scan(&question.ID, &question.CategoryID, &question.QuestionText, &question.CorrectAnswer); err != nil {
			return nil, err
		}
		questions = append(questions, question)
	}
	return questions, nil
}
