package usecase

import (
	"errors"
	"phishing-quest/adapter/repository"
	"phishing-quest/dto"
)

type GameUseCase struct {
	answerRepo    repository.IAnswerRepository
	userRepo      repository.IUserRepository
	userScoreRepo repository.IUserScoreRepository
}

func NewGameUseCase(answerRepo repository.IAnswerRepository, userRepo repository.IUserRepository, userScoreRepo repository.IUserScoreRepository) *GameUseCase {
	return &GameUseCase{
		answerRepo:    answerRepo,
		userRepo:      userRepo,
		userScoreRepo: userScoreRepo,
	}
}

func (guc *GameUseCase) ProcessAnswer(sbaDTO *dto.SubmitAnswerDTO) (*dto.AnswerResultDTO, error) {
	answer, err := guc.answerRepo.GetByID(sbaDTO.AnswerID)
	if err != nil {
		return nil, err
	}

	if answer.QuestionId != sbaDTO.QuestionID {
		return nil, errors.New("answer does not belong to the given question")
	}

	isCorrect := answer.IsCorrect
	if isCorrect {
		err = guc.userScoreRepo.IncrementScore(sbaDTO.UserID, 10) // Exemplo: 10 pontos por resposta correta
		if err != nil {
			return nil, err
		}
	}

	result := &dto.AnswerResultDTO{
		IsCorrect: isCorrect,
		Message:   "Resposta processada com sucesso!",
	}

	return result, nil
}
