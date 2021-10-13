package product

import (
	"sync"

	"github.com/google/uuid"
)

type memoryRepository struct {
	products map[uuid.UUID]Product
	sync.Mutex
}

func (r *memoryRepository) Fetch() []*Product {
	products := make([]*Product, 0)
	for _, product := range r.products {
		products = append(products, &product)
	}

	return products
}

func (r *memoryRepository) Get(id uuid.UUID) (*Product, error) {
	if product, ok := r.products[id]; ok {
		return &product, nil
	}
	return nil, ErrProductNotFound
}

func (r *memoryRepository) Add(product *Product) error {
	if _, ok := r.products[product.ID()]; ok {
		return ErrFailedToAddProduct
	}

	return r.save(product)
}

func (r *memoryRepository) Update(product *Product) error {
	if _, ok := r.products[product.ID()]; !ok {
		return ErrFailedToUpdateProduct
	}

	return r.save(product)
}

func (r *memoryRepository) save(product *Product) error {
	r.Lock()
	r.products[product.ID()] = *product
	r.Unlock()
	return nil
}

func (r *memoryRepository) Delete(id uuid.UUID) error {
	if _, ok := r.products[id]; !ok {
		return ErrFailedToDeleteProduct
	}

	delete(r.products, id)
	return nil
}

func NewMemoryRepository(dump ...Product) *memoryRepository {
	products := make(map[uuid.UUID]Product)
	for _, product := range dump {
		products[product.ID()] = product
	}

	return &memoryRepository{
		products: products,
	}
}
