package service

import (
	"test-backend/internal/models"
	"test-backend/internal/repository"
)

// UserService defines business logic for users.
type UserService interface {
	GetAll() []models.User
	GetByID(id int) (models.User, bool)
	Create(user models.User) models.User
	Update(id int, user models.User) (models.User, bool)
	Delete(id int) bool
}

// userService is a concrete implementation of UserService.
type userService struct {
	repo repository.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(r repository.UserRepository) UserService {
	return &userService{repo: r}
}

func (s *userService) GetAll() []models.User {
	return s.repo.GetAll()
}

func (s *userService) GetByID(id int) (models.User, bool) {
	return s.repo.GetByID(id)
}

func (s *userService) Create(user models.User) models.User {
	return s.repo.Create(user)
}

func (s *userService) Update(id int, user models.User) (models.User, bool) {
	return s.repo.Update(id, user)
}

func (s *userService) Delete(id int) bool {
	return s.repo.Delete(id)
}
