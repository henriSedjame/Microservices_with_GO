package models

import (
	"encoding/json"
	"io"
)

type Products []*Product

func (prods Products) ToJson(wr io.Writer) error {
	return json.NewEncoder(wr).Encode(prods)
}
