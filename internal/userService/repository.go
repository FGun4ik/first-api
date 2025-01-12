package userService

import (
	"FirstProject/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user models.User) (models.User, error)
	GetUser() ([]models.User, error)
	UpdateUserByID(id uint, user models.User) (models.User, error)
	DeleteUserById(id uint) error
	GetTasksForUser(id uint) ([]models.Task, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user models.User) (models.User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return models.User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetUser() ([]models.User, error) {
	var users []models.User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, user models.User) (models.User, error) {
	var existingUser models.User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return models.User{}, err
	}
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	r.db.Save(&existingUser)
	return existingUser, nil
}

func (r *userRepository) DeleteUserById(id uint) error {
	var user models.User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}
	r.db.Delete(&user)
	return nil
}

func (r *userRepository) GetTasksForUser(userID uint) ([]models.Task, error) {
	var tasks []models.Task
	err := r.db.Where("user_id = ?", userID).Find(&tasks).Error
	if err != nil {
		return nil, err
	}
	return tasks, nil
}
