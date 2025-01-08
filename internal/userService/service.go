package userService

type UserService struct {
	repo UserRepository
}

func NewUserService(repo UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(user User) (User, error) {
	return s.repo.CreateUser(user)
}

func (s *UserService) GetUser() ([]User, error) {
	return s.repo.GetUser()
}

func (s *UserService) UpdateUserById(id uint, user User) (User, error) {
	return s.repo.UpdateUserByID(id, user)
}

func (s *UserService) DeleteUserById(id uint) error {
	return s.repo.DeleteUserById(id)
}
