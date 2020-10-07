package test

import (
	"fmt"
	"github.com/hsedjame/products-api/src/main/models"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestValidation(t *testing.T) {
	product := models.Product{
		Name:  "Coffee",
		Price: 4.5,
		SKU:   "abc-ase-der",
	}

	err := product.Validate()

	assert.Nil(t, err, fmt.Sprintf("product is not valid : %s", err))

}
