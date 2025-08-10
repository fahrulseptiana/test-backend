package product

// Repository defines methods for product data access.
type Repository interface {
	GetAll() []Product
	GetByID(id int) (Product, bool)
	Create(product Product) Product
	Update(id int, product Product) (Product, bool)
	Delete(id int) bool
}

// InMemoryRepository is an in-memory implementation of Repository.
type InMemoryRepository struct {
	data   map[int]Product
	lastID int
}

// NewInMemoryRepository creates a new in-memory repository.
func NewInMemoryRepository() *InMemoryRepository {
	return &InMemoryRepository{data: make(map[int]Product)}
}

func (r *InMemoryRepository) GetAll() []Product {
	products := make([]Product, 0, len(r.data))
	for _, p := range r.data {
		products = append(products, p)
	}
	return products
}

func (r *InMemoryRepository) GetByID(id int) (Product, bool) {
	p, ok := r.data[id]
	return p, ok
}

func (r *InMemoryRepository) Create(product Product) Product {
	r.lastID++
	product.ID = r.lastID
	r.data[product.ID] = product
	return product
}

func (r *InMemoryRepository) Update(id int, product Product) (Product, bool) {
	if _, ok := r.data[id]; !ok {
		return Product{}, false
	}
	product.ID = id
	r.data[id] = product
	return product, true
}

func (r *InMemoryRepository) Delete(id int) bool {
	if _, ok := r.data[id]; !ok {
		return false
	}
	delete(r.data, id)
	return true
}
