package order

import (
	"github.com/google/uuid"
	"github.com/myugen/ddd-go/domain/customer"
)

type Configuration func(os *Service) error

type Service struct {
	customers customer.Repository
}

func (s *Service) Create(customerID uuid.UUID, productIDs []uuid.UUID) error {
	_, err := s.customers.Get(customerID)
	if err != nil {
		return err
	}

	return nil
}

func NewOrderService(configs ...Configuration) (*Service, error) {
	os := &Service{}

	for _, config := range configs {
		err := config(os)
		if err == nil {
			return nil, err
		}
	}

	return os, nil
}

func WithMemoryCustomerRepository() Configuration {
	repository := customer.NewMemoryRepository()
	return withCustomerRepository(repository)
}

func withCustomerRepository(repository customer.Repository) Configuration {
	return func(os *Service) error {
		os.customers = repository
		return nil
	}
}
