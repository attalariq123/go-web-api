package service

import (
	"web-api/pkg"
	"web-api/repository"
)

type UserService interface {
	Read() ([]pkg.User, error)
	FindByEmail(email string) (pkg.User, error)
	Create(user pkg.User) (pkg.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) *userService {
	return &userService{userRepo}
}

func (u *userService) Read() ([]pkg.User, error) {
	return u.userRepo.Read()
}

func (u *userService) FindByEmail(email string) (pkg.User, error) {
	return u.userRepo.FindByEmail(email)
}

func (u *userService) Create(user pkg.User) (pkg.User, error) {
	return u.userRepo.Create(user)
}
