package product_test

import (
	"testing"

	"github.com/myugen/ddd-go/domain/product"
	"github.com/stretchr/testify/suite"
)

type ProductSuite struct {
	suite.Suite
}

func (s *ProductSuite) TestNewProduct() {
	type testCase struct {
		test        string
		name        string
		description string
		price       float64
		expectedErr error
	}
	testCases := []testCase{
		{
			test:        "Empty values validation",
			name:        "",
			description: "test",
			price:       0,
			expectedErr: product.ErrInvalidItem,
		},
		{
			test:        "Valid item",
			name:        "test",
			description: "test",
			price:       4.5,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.test, func() {
			got, err := product.NewProduct(tc.name, tc.description, tc.price)
			if err == nil {
				s.NotNil(got, "product exists")
				s.Equal(tc.name, got.Name(), "product has correct name")
				s.Equal(tc.description, got.Description(), "product has correct description")
				s.Equal(tc.price, got.Price(), "product has correct price")
			}
			s.Equal(tc.expectedErr, err, "error validation")
		})
	}
}

func TestProductSuite(t *testing.T) {
	suite.Run(t, new(ProductSuite))
}
