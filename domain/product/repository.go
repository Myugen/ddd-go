package product

import (
	"errors"

	"github.com/google/uuid"
)

var (
	//ErrProductNotFound is returned when a product is not found.
	ErrProductNotFound = errors.New("the product was not found in the repository")
	// ErrFailedToAddProduct is returned when the product could not be added to the repository.
	ErrFailedToAddProduct = errors.New("failed to add the product to the repository")
	// ErrFailedToUpdateProduct is returned when the product could not be updated in the repository.
	ErrFailedToUpdateProduct = errors.New("failed to update the product in the repository")
	// ErrFailedToDeleteProduct is returned when the product could not be deleted in the repository.
	ErrFailedToDeleteProduct = errors.New("failed to delete the product in the repository")
)

type Repository interface {
	Fetch() []*Product
	Get(id uuid.UUID) (*Product, error)
	Add(product *Product) error
	Update(product *Product) error
	Delete(id uuid.UUID) error
}
