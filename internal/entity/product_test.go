package entity

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("product1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.Equal(t, "product1", product.Name)
	assert.Equal(t, 10, product.Price)
	assert.NotEmpty(t, product.ID)
	assert.NotEmpty(t, product.CreateAt)
}

func TestValidate(t *testing.T) {
	t.Run("ErrNameisRequired", func(t *testing.T) {
		_, err := NewProduct("", 10)

		assert.Equal(t, err, ErrNameisRequired)
	})

	t.Run("ErrPriceisRequired", func(t *testing.T) {
		_, err := NewProduct("product1", 0)
		expectedErr := errors.New("price is required")

		assert.Equal(t, err, expectedErr)
	})

	t.Run("ErrInvalidPrice", func(t *testing.T) {
		_, err := NewProduct("product1", -1)
		expectedErr := errors.New("invalid price")

		assert.Equal(t, err, expectedErr)
	})
}
