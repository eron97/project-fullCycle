package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewUser(t *testing.T) {
	result, err := NewUser("John Doe", "j@j.com", "1231989")
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result.ID)
	assert.NotEmpty(t, result.Password)
	assert.Equal(t, "John Doe", result.Name)
	assert.Equal(t, "j@j.com", result.Email)
}

func TestValidatePassword(t *testing.T) {
	result, err := NewUser("John Doe", "j@j.com", "1231989")
	assert.Nil(t, err)
	assert.True(t, result.ValidatePassword("1231989"))
	assert.False(t, result.ValidatePassword("123"))
	assert.NotEqual(t, "1231989", result.Password)
}
