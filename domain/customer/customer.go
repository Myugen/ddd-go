package customer

import (
	"errors"

	"github.com/google/uuid"

	tavern "github.com/myugen/ddd-go"
)

var (
	//ErrInvalidPerson is returned when the person is not valid
	ErrInvalidPerson = errors.New("a customer needs a valid person")
)

//Customer is an aggregate that combines all entities needed to represent a customer
type Customer struct {
	//person is the root entity of a customer
	//which means the person.ID is the main identifier for this aggregate
	person *tavern.Person
	// a customer can hold many products
	products []*tavern.Item
	// a customer can perform many transactions
	transaction []tavern.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &tavern.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:      person,
		products:    make([]*tavern.Item, 0),
		transaction: make([]tavern.Transaction, 0),
	}, nil
}
