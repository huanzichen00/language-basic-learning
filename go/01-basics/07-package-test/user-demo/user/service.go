package user

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) UserService {
	return UserService{
		repo: repo,
	}
}

func (s UserService) GetUserName(id int) (string, error) {
	user, err := s.repo.FindByID(id)
	if err != nil {
		return "", err
	}

	return user.Name, nil
}
