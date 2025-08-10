package repository

import "test-backend/internal/models"

// UserRepository defines methods for user data access.
type UserRepository interface {
	GetAll() []models.User
	GetByID(id int) (models.User, bool)
	Create(user models.User) models.User
	Update(id int, user models.User) (models.User, bool)
	Delete(id int) bool
}

// InMemoryUserRepository is an in-memory implementation of UserRepository.
type InMemoryUserRepository struct {
	data   map[int]models.User
	lastID int
}

// NewInMemoryUserRepository creates a new in-memory repository.
func NewInMemoryUserRepository() *InMemoryUserRepository {
	return &InMemoryUserRepository{data: make(map[int]models.User)}
}

func (r *InMemoryUserRepository) GetAll() []models.User {
	users := make([]models.User, 0, len(r.data))
	for _, u := range r.data {
		users = append(users, u)
	}
	return users
}

func (r *InMemoryUserRepository) GetByID(id int) (models.User, bool) {
	u, ok := r.data[id]
	return u, ok
}

func (r *InMemoryUserRepository) Create(user models.User) models.User {
	r.lastID++
	user.ID = r.lastID
	r.data[user.ID] = user
	return user
}

func (r *InMemoryUserRepository) Update(id int, user models.User) (models.User, bool) {
	if _, ok := r.data[id]; !ok {
		return models.User{}, false
	}
	user.ID = id
	r.data[id] = user
	return user, true
}

func (r *InMemoryUserRepository) Delete(id int) bool {
	if _, ok := r.data[id]; !ok {
		return false
	}
	delete(r.data, id)
	return true
}
