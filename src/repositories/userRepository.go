package repositories

import (
	"fmt"

	"finances/src/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *UserRepository {
	return &UserRepository{db}
}
func (u *UserRepository) Create(user *models.User) error {
	return u.db.Create(user).Error
}

func (u *UserRepository) Get(id string) (*models.User, error) {
	var user models.User
	err := u.db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (u *UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	err := u.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *UserRepository) Update(user *models.User) error {
	return u.db.Save(user).Error
}

func (u *UserRepository) Delete(id string) error {
	return u.db.Delete(&models.User{}, id).Error
}
