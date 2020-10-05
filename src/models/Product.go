package models

import (
	"encoding/json"
	"io"
)

type Product struct {
	ID int `json:"id"`
	Name string `json:"name"`
	Description string `json:"description"`
	Price float32 `json:"price"`
	SKU string `json:"sku"`
	CreationDate string `json:"-"`
	UpdateDate string `json:"-"`
	RemovalDate string `json:"-"`
}

func (p *Product) ToJson(wr io.Writer) error {
	return json.NewEncoder(wr).Encode(p)
}

func (p *Product) FromJson(rd io.Reader) error {
	return json.NewDecoder(rd).Decode(p)
}

type Products []*Product

func (prods Products) ToJson(wr io.Writer) error {
	return json.NewEncoder(wr).Encode(prods)
}
