package usecase

type UserService struct {
	repo domain.UserRepository
}

func NewUserService(repo domain.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(username, email, passwordHash string) (*domain.User, error) {
	user := &domain.User{
		Username:     username,
		Email:        email,
		PasswordHash: passwordHash,
	}
	err := s.repo.Save(user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (s *UserService) GetUserByID(id string) (*domain.User, error) {
	return s.repo.FindByID(id)
}
