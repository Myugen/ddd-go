package product

import (
	"errors"

	"github.com/google/uuid"

	tavern "github.com/myugen/ddd-go"
)

var (
	ErrInvalidItem = errors.New("a product needs a valid item")
)

//Product is an aggregate that combines all entities needed to represent a product
type Product struct {
	//item is the root entity of a product
	//which means the item.ID is the main identifier for this aggregate
	item  *tavern.Item
	price float64
	//quantity is the number of products in stock
	quantity int
}

func (p Product) ID() uuid.UUID {
	return p.item.ID
}

func (p *Product) WithID(id uuid.UUID) *Product {
	if p.item == nil {
		p.item = &tavern.Item{}
	}

	p.item.ID = id
	return p
}

func (p Product) Name() string {
	return p.item.Name
}

func (p *Product) WithName(name string) *Product {
	if p.item == nil {
		p.item = &tavern.Item{}
	}

	p.item.Name = name
	return p
}

func (p Product) Description() string {
	return p.item.Description
}

func (p *Product) WithDescription(description string) *Product {
	if p.item == nil {
		p.item = &tavern.Item{}
	}

	p.item.Description = description
	return p
}

func (p Product) Price() float64 {
	return p.price
}

func (p *Product) WithPrice(price float64) *Product {
	p.price = price
	return p
}

func NewProduct(name, description string, price float64) (*Product, error) {
	if name == "" || description == "" {
		return nil, ErrInvalidItem
	}

	item := &tavern.Item{
		ID:          uuid.New(),
		Name:        name,
		Description: description,
	}

	return &Product{
		item:     item,
		price:    price,
		quantity: 0,
	}, nil
}
