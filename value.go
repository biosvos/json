package json

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type Value interface {
	Get(paths ...string) Value
	AsSlice() []Value
	String() (string, error)
	Int() (int64, error)
	Bool() (bool, error)
	Bytes() []byte
}

func NewJsonValue(v []byte) Value {
	if v == nil {
		return NewErrorValue(errors.New("failed to new json value"))
	}

	m, err := tryMap(v)
	if err == nil {
		return NewMapValue(m)
	}

	s, err := trySlice(v)
	if err == nil {
		return NewSliceValue(s)
	}

	return NewPrimitiveValue(v)
}

func tryMap(v []byte) (map[string]any, error) {
	ret := map[string]any{}
	err := json.Unmarshal(v, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func trySlice(v []byte) ([]any, error) {
	var ret []any
	err := json.Unmarshal(v, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
