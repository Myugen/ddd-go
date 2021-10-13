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
	john, err := customer.NewCustomer("john")
	if err != nil {
		s.Fail(err.Error())
	}

	s.customers = customer.NewMemoryRepository(*john)
	s.existingCustomerID = john.ID()
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

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got, err := s.customers.Get(tc.id)
			if err == nil {
				s.NotNil(got, "customer exists")
				s.Equal(got.ID(), tc.id, "customer has correct id")
			} else {
				s.Equal(tc.expectedErr, err, "error validation")
			}
		})
	}

}

func (s *MemoryRepositorySuite) TestAddCustomer() {
	jane, err := customer.NewCustomer("jane")
	if err != nil {
		s.Fail(err.Error())
	}

	jacob, err := customer.NewCustomer("jacob")
	if err != nil {
		s.Fail(err.Error())
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

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			err := s.customers.Add(tc.customer)

			s.Equal(tc.expectedErr, err, "error on adding new customer")
		})
	}
}

func (s *MemoryRepositorySuite) TestUpdateCustomer() {
	jane, err := customer.NewCustomer("jane")
	if err != nil {
		s.Fail(err.Error())
	}

	jacob, err := customer.NewCustomer("jacob")
	if err != nil {
		s.Fail(err.Error())
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

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			err := s.customers.Update(tc.customer)

			s.Equal(tc.expectedErr, err, "error on adding new customer")
		})
	}
}

func TestMemoryRepositorySuite(t *testing.T) {
	suite.Run(t, new(MemoryRepositorySuite))
}
