package tax

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCalculateTaxAndSave(t *testing.T) {
	repository := &RepositoryMock{}
	repository.On("SaveTax", 5.0).Return(nil)
	repository.On("SaveTax", 0.0).Return(errors.New("amount is 0"))
	err := CalculateTaxAndSave(500.0, repository)
	assert.Nil(t, err)
	err = CalculateTaxAndSave(0.0, repository)
	assert.Error(t, err, "amount is 0")
	repository.AssertExpectations(t)
}
