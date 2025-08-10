package user

// Repository defines methods for user data access.
type Repository interface {
	GetAll() []User
	GetByID(id int) (User, bool)
	Create(user User) User
	Update(id int, user User) (User, bool)
	Delete(id int) bool
}

// InMemoryRepository is an in-memory implementation of Repository.
type InMemoryRepository struct {
	data   map[int]User
	lastID int
}

// NewInMemoryRepository creates a new in-memory repository.
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{data: make(map[int]User)}
}

func (r *InMemoryRepository) GetAll() []User {
	users := make([]User, 0, len(r.data))
	for _, u := range r.data {
		users = append(users, u)
	}
	return users
}

func (r *InMemoryRepository) GetByID(id int) (User, bool) {
	u, ok := r.data[id]
	return u, ok
}

func (r *InMemoryRepository) Create(user User) User {
	r.lastID++
	user.ID = r.lastID
	r.data[user.ID] = user
	return user
}

func (r *InMemoryRepository) Update(id int, user User) (User, bool) {
	if _, ok := r.data[id]; !ok {
		return User{}, false
	}
	user.ID = id
	r.data[id] = user
	return user, true
}

func (r *InMemoryRepository) Delete(id int) bool {
	if _, ok := r.data[id]; !ok {
		return false
	}
	delete(r.data, id)
	return true
}
