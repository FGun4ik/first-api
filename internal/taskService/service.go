package taskService

import "fmt"

type TaskService struct {
	repo TaskRepository
}

func NewService(repo TaskRepository) *TaskService {
	return &TaskService{repo: repo}
}

func (s *TaskService) CreateTask(task Task) (Task, error) {
	if task.UserID == 0 {
		return Task{}, fmt.Errorf("user_id is required")
	}
	return s.repo.CreateTask(task)
}

func (s *TaskService) GetAllTasks() ([]Task, error) {
	return s.repo.GetAllTasks()

}

func (s *TaskService) UpdateTaskById(id uint, task Task) (Task, error) {
	return s.repo.UpdateTaskByID(id, task)
}

func (s *TaskService) DeleteTaskById(id uint) error {
	return s.repo.DeleteTaskByID(id)
}

func (s *TaskService) GetTasksByUserID(userID uint) ([]Task, error) {
	return s.repo.GetTasksByUserID(userID)
}
