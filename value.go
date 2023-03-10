package json

import "github.com/pkg/errors"

type Value interface {
	Get(paths ...string) Value
	Slice() []Value
	String() string
	Err() error
}

func NewJsonValue(v []byte) Value {
	if v == nil {
		return newErrorValue(errors.New("failed to new json value. cause: v is nil"))
	}

	m, err := toMap(v)
	if err == nil {
		return newMapValue(m)
	}

	s, err := toSlice(v)
	if err == nil {
		return newSliceValue(s)
	}

	return newPrimitiveValue(v)
}
