package order

import (
	"fmt"
	"log"

	"github.com/google/uuid"
	"github.com/myugen/ddd-go/domain/customer"
	"github.com/myugen/ddd-go/domain/product"
)

type Configuration func(os *Service) error

type Service struct {
	customers customer.Repository
	products  product.Repository
}

//Create will chain together all repositories to create an order for a customer
//will return the collected price of all Products
func (s *Service) Create(customerID uuid.UUID, productIDs []uuid.UUID) (float64, error) {
	c, err := s.customers.Get(customerID)
	if err != nil {
		return 0, err
	}

	products := make([]product.Product, 0)
	total := float64(0)
	for _, productID := range productIDs {
		p, err := s.products.Get(productID)
		if err != nil {
			return 0, err
		}
		products = append(products, *p)
		total += p.Price()
	}

	log.Println(fmt.Sprintf("Customer: %s has ordered %d products", c.Name(), len(products)))

	return total, nil
}

func NewService(configs ...Configuration) (*Service, error) {
	os := new(Service)

	for _, config := range configs {
		err := config(os)
		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func WithMemoryCustomerRepository(customers ...customer.Customer) Configuration {
	repository := customer.NewMemoryRepository(customers...)
	return withCustomerRepository(repository)
}

func withCustomerRepository(repository customer.Repository) Configuration {
	return func(os *Service) error {
		os.customers = repository
		return nil
	}
}

func WithMemoryProductRepository(products ...product.Product) Configuration {
	repository := product.NewMemoryRepository(products...)
	return withProductRepository(repository)
}

func withProductRepository(repository product.Repository) Configuration {
	return func(os *Service) error {
		os.products = repository
		return nil
	}
}
