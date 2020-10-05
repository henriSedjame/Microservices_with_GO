package core

import "io"

type Serializable interface {
	ToJson(wr io.Writer) error
	FromJson(rd io.Reader) error
}
