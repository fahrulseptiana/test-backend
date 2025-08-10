package user

// Service defines business logic for users.
type Service interface {
	GetAll() []User
	GetByID(id int) (User, bool)
	Create(user User) User
	Update(id int, user User) (User, bool)
	Delete(id int) bool
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

func (s *service) Create(user User) User {
	return s.repo.Create(user)
}

func (s *service) Update(id int, user User) (User, bool) {
	return s.repo.Update(id, user)
}

func (s *service) Delete(id int) bool {
	return s.repo.Delete(id)
}
