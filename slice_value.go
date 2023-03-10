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

func (s *sliceValue) Get(paths ...string) Value {
	if len(paths) == 0 {
		return s
	}
	return newErrorValue(errors.New("failed to get"))
}

func (s *sliceValue) Slice() []Value {
	var ret []Value
	for _, item := range s.value {
		marshal, _ := json.Marshal(item)
		value := NewJsonValue(marshal)
		ret = append(ret, value)
	}
	return ret
}

func (s *sliceValue) String() string {
	marshal, _ := json.Marshal(s.value)
	return string(marshal)
}

func (s *sliceValue) Err() error {
	return nil
}
