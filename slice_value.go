package json

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ Value = &sliceValue{}

type sliceValue struct {
	value []any
}

func newSliceValue(value []any) *sliceValue {
	return &sliceValue{value: value}
}

func (s *sliceValue) Get(paths ...string) (Value, error) {
	if len(paths) == 0 {
		return s, nil
	}
	return nil, errors.New("failed to get")
}

func (s *sliceValue) List(paths ...string) ([]Value, error) {
	return slice(s, paths...)
}

func (s *sliceValue) slice() ([]Value, error) {
	var ret []Value
	for _, item := range s.value {
		marshal, _ := json.Marshal(item)
		value, err := NewJsonValue(marshal)
		if err != nil {
			return nil, errors.WithStack(err)
		}
		ret = append(ret, value)
	}
	return ret, nil
}

func (s *sliceValue) String() string {
	return string(s.Bytes())
}

func (s *sliceValue) Bytes() []byte {
	ret, _ := json.Marshal(s.value)
	return ret
}
