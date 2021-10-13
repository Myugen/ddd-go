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

func (c *Customer) ID() uuid.UUID {
	if c.person == nil {
		c.person = &tavern.Person{}
	}

	return c.person.ID
}

func (c *Customer) WithID(id uuid.UUID) *Customer {
	if c.person == nil {
		c.person = &tavern.Person{}
	}

	c.person.ID = id

	return c
}

func (c *Customer) Name() string {
	if c.person == nil {
		c.person = &tavern.Person{}
	}

	return c.person.Name
}

func (c *Customer) WithName(name string) *Customer {
	if c.person == nil {
		c.person = &tavern.Person{}
	}

	c.person.Name = name

	return c
}

func NewCustomer(name string) (*Customer, error) {
	if name == "" {
		return nil, ErrInvalidPerson
	}

	person := &tavern.Person{
		ID:   uuid.New(),
		Name: name,
	}

	customer := &Customer{
		person:      person,
		products:    make([]*tavern.Item, 0),
		transaction: make([]tavern.Transaction, 0),
	}
	return customer, nil
}
