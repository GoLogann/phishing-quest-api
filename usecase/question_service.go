package usecase


type QuestionService struct {
    repo domain.QuestionRepository
}

func NewQuestionService(repo domain.QuestionRepository) *QuestionService {
    return &QuestionService{repo: repo}
}

func (s *QuestionService) CreateQuestion(categoryID int, questionText, correctAnswer string) (*domain.Question, error) {
    question := &domain.Question{
        CategoryID:   categoryID,
        QuestionText: questionText,
        CorrectAnswer: correctAnswer,
    }
    err := s.repo.Save(question)
    if err != nil {
        return nil, err
    }
    return question, nil
}

func (s *QuestionService) GetQuestionByID(id int) (*domain.Question, error) {
    return s.repo.FindByID(id)
}

func (s *QuestionService) GetQuestionsByCategory(categoryID int) ([]*domain.Question, error) {
    return s.repo.FindByCategoryID(categoryID)
}