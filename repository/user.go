package repository

import (
	"web-api/pkg"

	"gorm.io/gorm"
)

type UserRepository interface {
	Read() ([]pkg.User, error)
	FindByEmail(email string) (pkg.User, error)
	Create(user pkg.User) (pkg.User, error)
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

func (r *userRepository) Read() ([]pkg.User, error) {
	var users []pkg.User

	err := r.db.Find(&users).Error

	return users, err
}

func (r *userRepository) FindByEmail(email string) (pkg.User, error) {
	var user pkg.User

	err := r.db.Where("email = ?", email).First(&user).Error

	return user, err
}

func (r *userRepository) Create(user pkg.User) (pkg.User, error) {

	err := r.db.Create(&user).Error

	return user, err
}
