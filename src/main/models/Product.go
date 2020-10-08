package models

import (
	"github.com/go-playground/validator"
	"github.com/hsedjame/products-api/src/main/core"
)

// Product represents the product for this application
//
// swagger:model
type Product struct {
	// the id for this product
	//
	// required: true
	ID int `json:"id"`

	// the name for this product
	//
	// required: true
	Name string `json:"name" validate:"required"`

	// the description for this product
	//
	// required: false
	Description string `json:"description"`

	// the price for this product
	//
	// min: 0
	Price float32 `json:"price" validate:"gt=0"`

	// the sku for this product
	//
	// required: true
	// pattern: [a-z]+-[a-z]+-[a-z]+
	SKU          string `json:"sku" validate:"required,sku"`
	CreationDate string `json:"-"`
	UpdateDate   string `json:"-"`
	RemovalDate  string `json:"-"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("sku", core.SkuValidation)
	if err != nil {
		return err
	}
	return validate.Struct(p)
}
