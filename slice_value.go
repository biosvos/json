package json

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ Value = &SliceValue{}

type SliceValue struct {
	value []any
}

func NewSliceValue(value []any) *SliceValue {
	return &SliceValue{value: value}
}

func (s *SliceValue) Get(paths ...string) Value {
	if len(paths) == 0 {
		return s
	}
	return NewErrorValue(errors.Errorf("unknown path %v", paths))
}

func (s *SliceValue) AsSlice() []Value {
	var ret []Value
	for _, v := range s.value {
		marshal, err := json.Marshal(v)
		if err != nil {
			panic(err)
		}
		ret = append(ret, NewJsonValue(marshal))
	}
	return ret
}

func (s *SliceValue) String() (string, error) {
	return string(s.Bytes()), nil
}

func (s *SliceValue) Int() (int64, error) {
	return 0, errors.New("failed to convert")
}

func (s *SliceValue) Bool() (bool, error) {
	return false, errors.New("failed to convert")
}

func (s *SliceValue) Bytes() []byte {
	ret, err := json.Marshal(s.value)
	if err != nil {
		panic(err)
	}
	return ret
}
