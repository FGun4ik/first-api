package taskServise

type TaskService struct {
	repo MessageRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks(task Task) (Task, error) {
	return s.repo.GetAllTasks
}
