package customer_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/myugen/ddd-go/domain/customer"
)

func TestNewCustomer(t *testing.T) {
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
		t.Run(tt.test, func(t *testing.T) {
			got, err := customer.NewCustomer(tt.name)
			if err == nil {
				assert.NotNil(t, got, "customer exists")
				assert.Equal(t, got.Name(), tt.name, "customer has correct name")
			} else {
				assert.Equal(t, tt.expectedErr, err, "error validation")
			}
		})
	}
}
