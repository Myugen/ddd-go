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

	for _, tc := range testCases {
		s.Run(tc.test, func() {
			got, err := customer.NewCustomer(tc.name)
			if err == nil {
				s.NotNil(got, "customer exists")
				s.Equal(got.Name(), tc.name, "customer has correct name")
			}

			s.Equal(tc.expectedErr, err, "error validation")
		})
	}
}

func TestCustomerTestSuite(t *testing.T) {
	suite.Run(t, new(CustomerSuite))
}
