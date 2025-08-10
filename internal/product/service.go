package product

// Service defines business logic for products.
type Service interface {
	GetAll() []Product
	GetByID(id int) (Product, bool)
	Create(product Product) Product
	Update(id int, product Product) (Product, bool)
	Delete(id int) bool
}

type service struct {
	repo Repository
}

// NewService creates a new Service.
func NewService(r Repository) Service {
	return &service{repo: r}
}

func (s *service) GetAll() []Product {
	return s.repo.GetAll()
}

func (s *service) GetByID(id int) (Product, bool) {
	return s.repo.GetByID(id)
}

func (s *service) Create(product Product) Product {
	return s.repo.Create(product)
}

func (s *service) Update(id int, product Product) (Product, bool) {
	return s.repo.Update(id, product)
}

func (s *service) Delete(id int) bool {
	return s.repo.Delete(id)
}
