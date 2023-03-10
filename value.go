package json

import "github.com/pkg/errors"

type Value interface {
	Get(paths ...string) (Value, error)
	AsSlice() ([]Value, error)
	String() string
	Bytes() []byte
}

func NewJsonValue(v []byte) (Value, error) {
	if v == nil {
		return nil, errors.New("failed to new json value. cause: v is nil")
	}

	bytes := byteList(v)
	m, err := bytes.toMap()
	if err == nil {
		return newMapValue(m), nil
	}

	s, err := bytes.toSlice()
	if err == nil {
		return newSliceValue(s), nil
	}

	return newPrimitiveValue(v), nil
}
