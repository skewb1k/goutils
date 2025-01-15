package optional

import (
	"encoding/json"
)

// Field is a wrapper for a value that may or may not be defined.
type Field[T any] struct {
	Defined bool
	Value   *T
}

func (o *Field[T]) UnmarshalJSON(data []byte) error {
	o.Defined = true
	return json.Unmarshal(data, &o.Value)
}
