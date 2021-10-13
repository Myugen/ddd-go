package order_test

import (
	"testing"

	"github.com/google/uuid"

	"github.com/myugen/ddd-go/domain/customer"

	"github.com/myugen/ddd-go/domain/product"

	"github.com/myugen/ddd-go/services/order"
	"github.com/stretchr/testify/suite"
)

type OrderSuite struct {
	suite.Suite
	service    *order.Service
	products   []product.Product
	customerID uuid.UUID
}

func (s *OrderSuite) initProducts() []product.Product {
	beer, err := product.NewProduct("Beer", "A pint of beer", 1.99)
	if err != nil {
		s.Fail(err.Error())
	}

	peanuts, err := product.NewProduct("Peanuts", "A health snack", 0.99)
	if err != nil {
		s.Fail(err.Error())
	}

	wine, err := product.NewProduct("Wine", "A glass of wine", 2.49)
	if err != nil {
		s.Fail(err.Error())
	}

	products := []product.Product{
		*beer, *peanuts, *wine,
	}
	return products
}

func (s *OrderSuite) initCustomer() customer.Customer {
	john, err := customer.NewCustomer("John")
	if err != nil {
		s.Fail(err.Error())
	}
	return *john
}

func (s *OrderSuite) SetupTest() {
	c := s.initCustomer()
	products := s.initProducts()
	service, err := order.NewService(
		order.WithMemoryCustomerRepository(c),
		order.WithMemoryProductRepository(products...),
	)
	if err != nil {
		s.Fail(err.Error())
	}

	s.service = service
	s.products = products
	s.customerID = c.ID()
}

func (s *OrderSuite) TestCreateOrder() {
	type testCase struct {
		name          string
		customerID    uuid.UUID
		productIDs    []uuid.UUID
		expectedPrice float64
		expectedErr   error
	}

	validProducts := []product.Product{
		//2 beers
		s.products[0],
		s.products[0],
		// 1 wine
		s.products[2],
		// 3 peanuts
		s.products[1],
		s.products[1],
		s.products[1],
	}
	validProductIDs := make([]uuid.UUID, 0)
	validProductsPrice := float64(0)

	for _, p := range validProducts {
		validProductIDs = append(validProductIDs, p.ID())
		validProductsPrice += p.Price()
	}

	testCases := []testCase{
		{
			name:          "Valid order",
			customerID:    s.customerID,
			productIDs:    validProductIDs,
			expectedPrice: validProductsPrice,
			expectedErr:   nil,
		},
		{
			name:          "Invalid customer in order",
			customerID:    uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			productIDs:    validProductIDs,
			expectedPrice: 0,
			expectedErr:   customer.ErrCustomerNotFound,
		},
		{
			name:          "Invalid product in order",
			customerID:    s.customerID,
			productIDs:    []uuid.UUID{uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479")},
			expectedPrice: 0,
			expectedErr:   product.ErrProductNotFound,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got, err := s.service.Create(tc.customerID, tc.productIDs)
			s.Equal(got, tc.expectedPrice)
			s.Equal(err, tc.expectedErr)
		})
	}
}

func TestOrderSuite(t *testing.T) {
	suite.Run(t, new(OrderSuite))
}
