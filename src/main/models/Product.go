package models

import (
	"encoding/json"
	"github.com/go-playground/validator"
	"github.com/hsedjame/products-api/src/main/core"
	"io"
)

type Product struct {
	ID           int     `json:"id"`
	Name         string  `json:"name" validate:"required"`
	Description  string  `json:"description"`
	Price        float32 `json:"price" validate:"gt=0"`
	SKU          string  `json:"sku" validate:"required,sku"`
	CreationDate string  `json:"-"`
	UpdateDate   string  `json:"-"`
	RemovalDate  string  `json:"-"`
}

func (p *Product) Validate() error {
	validate := validator.New()
	err := validate.RegisterValidation("sku", core.SkuValidation)
	if err != nil {
		return err
	}
	return validate.Struct(p)
}

func (p *Product) ToJson(wr io.Writer) error {
	return json.NewEncoder(wr).Encode(p)
}

func (p *Product) FromJson(rd io.Reader) error {
	return json.NewDecoder(rd).Decode(p)
}
