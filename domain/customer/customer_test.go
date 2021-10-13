package customer_test

import (
	"testing"

	"github.com/myugen/ddd-go/domain/customer"
	"github.com/stretchr/testify/suite"
)

type CustomerSuite struct {
	suite.Suite
}

func (s *CustomerSuite) TestNewCustomer() {
	type testCase struct {
		test        string
		name        string
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "Empty name validation",
			name:        "",
			expectedErr: customer.ErrInvalidPerson,
		},
		{
			test:        "Valid name",
			name:        "John",
			expectedErr: nil,
		},
	}

	for _, tt := range testCases {
		s.Run(tt.test, func() {
			got, err := customer.NewCustomer(tt.name)
			if err == nil {
				s.NotNil(got, "customer exists")
				s.Equal(got.Name(), tt.name, "customer has correct name")
			}

			s.Equal(tt.expectedErr, err, "error validation")
		})
	}
}

func TestCustomerTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerSuite))
}
