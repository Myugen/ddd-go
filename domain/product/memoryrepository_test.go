package product_test

import (
	"testing"

	"github.com/myugen/ddd-go/domain/product"

	"github.com/google/uuid"
	"github.com/stretchr/testify/suite"
)

type MemoryRepositorySuite struct {
	suite.Suite
	products          product.Repository
	existingProductID uuid.UUID
}

func (s *MemoryRepositorySuite) SetupTest() {
	beer, err := product.NewProduct("beer", "a pint of beer", 3.95)
	if err != nil {
		s.Fail(err.Error())
	}

	s.products = product.NewMemoryRepository(*beer)
	s.existingProductID = beer.ID()
}

func (s *MemoryRepositorySuite) TestFetchProducts() {
	expected := []uuid.UUID{s.existingProductID}
	type testCase struct {
		name        string
		expectedIDs []uuid.UUID
		expectedLen int
	}

	testCases := []testCase{
		{
			name:        "Fetch products",
			expectedIDs: expected,
			expectedLen: len(expected),
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got := s.products.Fetch()
			s.Len(got, tc.expectedLen, "fetched products has correct length")
			s.Condition(
				func() (success bool) {
					for _, expectedID := range tc.expectedIDs {
						for _, p := range got {
							if p.ID() != expectedID {
								return false
							}
						}
					}

					return true
				},
				"fetched products has expected ones")
		})
	}
}

func (s *MemoryRepositorySuite) TestGetProduct() {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Product not found by ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: product.ErrProductNotFound,
		},
		{

			name:        "Product found by ID",
			id:          s.existingProductID,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			got, err := s.products.Get(tc.id)
			if err == nil {
				s.NotNil(got, "product exists")
				s.Equal(got.ID(), tc.id, "product has correct id")
			} else {
				s.Equal(tc.expectedErr, err, "error validation")
			}
		})
	}

}

func (s *MemoryRepositorySuite) TestAddProduct() {
	mead, err := product.NewProduct("mead", "a pint of mead", 5.50)
	if err != nil {
		s.Fail(err.Error())
	}

	wine, err := product.NewProduct("wine", "a glass of wine", 4.50)
	if err != nil {
		s.Fail(err.Error())
	}

	type testCase struct {
		name        string
		product     *product.Product
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Valid product",
			product:     mead,
			expectedErr: nil,
		},
		{
			name:        "Not valid product",
			product:     wine.WithID(s.existingProductID),
			expectedErr: product.ErrFailedToAddProduct,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			err := s.products.Add(tc.product)
			s.Equal(tc.expectedErr, err, "error on adding new product")
		})
	}
}

func (s *MemoryRepositorySuite) TestUpdateProduct() {
	mead, err := product.NewProduct("mead", "a pint of mead", 5.50)
	if err != nil {
		s.Fail(err.Error())
	}

	wine, err := product.NewProduct("wine", "a glass of wine", 4.50)
	if err != nil {
		s.Fail(err.Error())
	}

	type testCase struct {
		name        string
		product     *product.Product
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Valid product",
			product:     wine.WithID(s.existingProductID),
			expectedErr: nil,
		},
		{
			name:        "Not valid product",
			product:     mead,
			expectedErr: product.ErrFailedToUpdateProduct,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			err := s.products.Update(tc.product)
			s.Equal(tc.expectedErr, err, "error on adding new product")
		})
	}
}

func (s *MemoryRepositorySuite) TestDeleteProduct() {
	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}

	testCases := []testCase{
		{
			name:        "Not valid product ID",
			id:          uuid.MustParse("f47ac10b-58cc-0372-8567-0e02b2c3d479"),
			expectedErr: product.ErrFailedToDeleteProduct,
		},
		{
			name:        "Valid product ID",
			id:          s.existingProductID,
			expectedErr: nil,
		},
	}

	for _, tc := range testCases {
		s.Run(tc.name, func() {
			err := s.products.Delete(tc.id)
			s.Equal(tc.expectedErr, err, "error on adding new product")
		})
	}
}

func TestMemoryRepositorySuite(t *testing.T) {
	suite.Run(t, new(MemoryRepositorySuite))
}
