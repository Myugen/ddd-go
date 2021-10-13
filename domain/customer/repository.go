package customer

import (
	"errors"

	"github.com/google/uuid"
)

var (
	// ErrCustomerNotFound is returned when a customer is not found.
	ErrCustomerNotFound = errors.New("the customer was not found in the repository")
	// ErrFailedToAddCustomer is returned when the customer could not be added to the repository.
	ErrFailedToAddCustomer = errors.New("failed to add the customer to the repository")
	// ErrFailedToUpdateCustomer is returned when the customer could not be updated in the repository.
	ErrFailedToUpdateCustomer = errors.New("failed to update the customer in the repository")
)

type Repository interface {
	Get(id uuid.UUID) (*Customer, error)
	Add(customer *Customer) error
	Update(customer *Customer) error
}
