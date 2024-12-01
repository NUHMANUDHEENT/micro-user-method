package repo

import (
	"fmt"
	"golang-application/internal/model"

	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUser(user model.User) (int, error)
	GetUser(id int) (*model.User, error)
	UpdateUser(user model.User) error
	DeleteUser(id int) error
	ListUser() ([]model.User, error)
	ListUserNames() ([]string, error)
}
type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}
func (r *userRepository) CreateUser(user model.User) (int, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return 0, fmt.Errorf("error creating user: %v", err)
	}
	if err := r.db.First(&user, "email=?", user.Email).Error; err != nil {
		return 0, fmt.Errorf("error creating user: %v", err)
	}

	return int(user.ID), nil
}
func (r *userRepository) GetUser(id int) (*model.User, error) {
	var user model.User
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("user not found")
	}
	if err := r.db.First(&user, id).Error; err != nil {
		return nil, fmt.Errorf("error getting user: %v", err)
	}
	return &user, nil
}
func (r *userRepository) UpdateUser(user model.User) error {
	var userData model.User
	if err := r.db.First(&userData, user.ID).Error; err != nil {
		return fmt.Errorf("user not found")
	}
	if err := r.db.Model(&user).Updates(user).Error; err != nil {
		return fmt.Errorf("error updating user: %v", err)
	}
	return nil

}
func (r *userRepository) DeleteUser(id int) error {
	var userData model.User
	if err := r.db.First(&userData, id).Error; err != nil {
		return fmt.Errorf("user not found")
	}
	if err := r.db.Delete(&userData).Error; err != nil {
		return fmt.Errorf("error deleting user: %v", err)
	}
	return nil
}
func (r *userRepository) ListUser() ([]model.User, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("error getting all users: %v", err)
	}
	return users, nil
}
func (r *userRepository) ListUserNames() ([]string, error) {
	var users []model.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, fmt.Errorf("error getting all users: %v", err)
	}
	var userNames []string
	for _, user := range users {
		userNames = append(userNames, user.Name)
	}
	return userNames, nil
}
