package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetUser() ([]User, error)
	UpdateUserByID(id uint, user User) (User, error)
	DeleteUserById(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user User) (User, error) {
	result := r.db.Create(&user)
	if result.Error != nil {
		return User{}, result.Error
	}
	return user, nil
}

func (r *userRepository) GetUser() ([]User, error) {
	var users []User
	err := r.db.Find(&users).Error
	return users, err
}

func (r *userRepository) UpdateUserByID(id uint, user User) (User, error) {
	var existingUser User
	if err := r.db.First(&existingUser, id).Error; err != nil {
		return User{}, err
	}
	existingUser.Email = user.Email
	existingUser.Password = user.Password
	r.db.Save(&existingUser)
	return existingUser, nil
}

func (r *userRepository) DeleteUserById(id uint) error {
	var user User
	if err := r.db.First(&user, id).Error; err != nil {
		return err
	}
	r.db.Delete(&user)
	return nil
}
