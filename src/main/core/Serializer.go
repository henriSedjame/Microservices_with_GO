package core

import (
	"encoding/json"
	"io"
)

func ToJson(i interface{}, wr io.Writer) error {
	return json.NewEncoder(wr).Encode(i)
}

func FromJson(i interface{}, rd io.Reader) error {
	return json.NewDecoder(rd).Decode(i)
}
