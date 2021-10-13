package customer

import (
	"sync"

	"github.com/google/uuid"
)

type memoryRepository struct {
	customers map[uuid.UUID]Customer
	sync.Mutex
}

func (r *memoryRepository) Get(id uuid.UUID) (*Customer, error) {
	if customer, ok := r.customers[id]; ok {
		return &customer, nil
	}
	return nil, ErrCustomerNotFound
}

func (r *memoryRepository) Add(customer *Customer) error {
	if _, ok := r.customers[customer.ID()]; ok {
		return ErrFailedToAddCustomer
	}

	return r.save(customer)
}

func (r *memoryRepository) Update(customer *Customer) error {
	if _, ok := r.customers[customer.ID()]; !ok {
		return ErrFailedToUpdateCustomer
	}

	return r.save(customer)
}

func (r *memoryRepository) save(customer *Customer) error {
	r.Lock()
	r.customers[customer.ID()] = *customer
	r.Unlock()
	return nil
}

func NewMemoryRepository() *memoryRepository {
	return &memoryRepository{
		customers: make(map[uuid.UUID]Customer),
	}
}
