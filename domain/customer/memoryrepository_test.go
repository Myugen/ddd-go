package customer_test

import (
	"testing"

	"github.com/google/uuid"
	"github.com/myugen/ddd-go/domain/customer"
	"github.com/stretchr/testify/suite"
)

type MemoryRepositorySuite struct {
	suite.Suite
	customers          customer.Repository
	existingCustomerID uuid.UUID
}

func (s *MemoryRepositorySuite) SetupTest() {
	c, err := customer.NewCustomer("john")
	if err != nil {
		s.Error(err)
	}

	s.customers = customer.NewMemoryRepository(*c)
	s.existingCustomerID = c.ID()
}

func (s *MemoryRepositorySuite) TestGetCustomer() {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Customer not found by ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{

			name:        "Customer found by ID",
			id:          s.existingCustomerID,
			expectedErr: nil,
		},
	}

	for _, tt := range testCases {
		s.Run(tt.name, func() {
			got, err := s.customers.Get(tt.id)
			if err == nil {
				s.NotNil(got, "customer exists")
				s.Equal(got.ID(), tt.id, "customer has correct id")
			} else {
				s.Equal(tt.expectedErr, err, "error validation")
			}
		})
	}

}

func (s *MemoryRepositorySuite) TestAddCustomer() {
	jane, err := customer.NewCustomer("jane")
	if err != nil {
		s.Error(err)
	}

	jacob, err := customer.NewCustomer("jacob")
	if err != nil {
		s.Error(err)
	}

	type testCase struct {
		name        string
		customer    *customer.Customer
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Valid customer",
			customer:    jane,
			expectedErr: nil,
		},
		{
			name:        "Not valid customer",
			customer:    jacob.WithID(s.existingCustomerID),
			expectedErr: customer.ErrFailedToAddCustomer,
		},
	}

	for _, tt := range testCases {
		s.Run(tt.name, func() {
			err := s.customers.Add(tt.customer)

			s.Equal(tt.expectedErr, err, "error on adding new customer")
		})
	}
}

func (s *MemoryRepositorySuite) TestUpdateCustomer() {
	jane, err := customer.NewCustomer("jane")
	if err != nil {
		s.Error(err)
	}

	jacob, err := customer.NewCustomer("jacob")
	if err != nil {
		s.Error(err)
	}

	type testCase struct {
		name        string
		customer    *customer.Customer
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Valid customer",
			customer:    jacob.WithID(s.existingCustomerID),
			expectedErr: nil,
		},
		{
			name:        "Not valid customer",
			customer:    jane,
			expectedErr: customer.ErrFailedToUpdateCustomer,
		},
	}

	for _, tt := range testCases {
		s.Run(tt.name, func() {
			err := s.customers.Update(tt.customer)

			s.Equal(tt.expectedErr, err, "error on adding new customer")
		})
	}
}

func TestMemoryRepositorySuite(t *testing.T) {
	suite.Run(t, new(MemoryRepositorySuite))
}
