package json

import "github.com/pkg/errors"

type Value interface {
	Get(paths ...string) (Value, error)
	List(paths ...string) ([]Value, error)
	slice() ([]Value, error)
	String() string
	Bytes() []byte
}

func slice(v Value, paths ...string) ([]Value, error) {
	get, err := v.Get(paths...)
	if err != nil {
		return nil, err
	}
	return get.slice()
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
