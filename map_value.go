package json

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
)

var _ Value = &MapValue{}

type MapValue struct {
	value map[string]any
}

func (m *MapValue) Get(paths ...string) Value {
	if len(paths) == 0 {
		return m
	}
	marshal, err := json.Marshal(m.value[paths[0]])
	if err != nil {
		panic(err)
	}
	value := NewJsonValue(marshal)
	return value.Get(paths[1:]...)
}

func NewMapValue(value map[string]any) *MapValue {
	return &MapValue{
		value: value,
	}
}

func (m *MapValue) get(path string) Value {
	marshal, err := json.Marshal(m.value[path])
	if err != nil {
		panic(err)
	}
	return NewJsonValue(marshal)
}

func (m *MapValue) Bytes() []byte {
	ret, err := json.Marshal(m.value)
	if err != nil {
		panic(err)
	}
	return ret
}

func (m *MapValue) AsSlice() []Value {
	slice, err := convert[[]map[string]any](m.value)
	if err != nil {
		panic(err)
	}
	var ret []Value
	for _, item := range slice {
		marshal, err := json.Marshal(item)
		if err != nil {
			panic(err)
		}
		ret = append(ret, NewJsonValue(marshal))
	}
	return ret
}

func (m *MapValue) String() string {
	return string(m.Bytes())
}

func (m *MapValue) Int() (int64, error) {
	str := m.String()
	ret, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return ret, nil
}

func (m *MapValue) Bool() (bool, error) {
	str := m.String()
	ret, err := strconv.ParseBool(str)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return ret, nil
}
