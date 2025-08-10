package user

import "golang.org/x/crypto/bcrypt"

// Service defines business logic for users.
type Service interface {
	GetAll() []User
	GetByID(id int) (User, bool)
	GetByEmail(email string) (User, bool)
	Create(user User) User
	Update(id int, user User) (User, bool)
	Delete(id int) bool
	Authenticate(email, password string) (User, bool)
}

// service is a concrete implementation of Service.
type service struct {
	repo Repository
}

// NewService creates a new Service.
func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() []User {
	return s.repo.GetAll()
}

func (s *service) GetByID(id int) (User, bool) {
	return s.repo.GetByID(id)
}

func (s *service) GetByEmail(email string) (User, bool) {
	return s.repo.GetByEmail(email)
}

func (s *service) Create(user User) User {
	if user.Password != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashed)
	}
	return s.repo.Create(user)
}

func (s *service) Update(id int, user User) (User, bool) {
	if user.Password != "" {
		hashed, _ := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		user.Password = string(hashed)
	}
	return s.repo.Update(id, user)
}

func (s *service) Delete(id int) bool {
	return s.repo.Delete(id)
}

func (s *service) Authenticate(email, password string) (User, bool) {
	user, ok := s.repo.GetByEmail(email)
	if !ok {
		return User{}, false
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return User{}, false
	}
	return user, true
}
