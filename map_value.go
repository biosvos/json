package json

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ Value = &mapValue{}

type mapValue struct {
	value map[string]any
}

func newMapValue(value map[string]any) *mapValue {
	return &mapValue{
		value: value,
	}
}

func (m *mapValue) Get(paths ...string) Value {
	if len(paths) == 0 {
		return m
	}
	v, ok := m.value[paths[0]]
	if !ok {
		return newErrorValue(errors.Errorf("not found path %v", paths[0]))
	}
	marshal, _ := json.Marshal(v)
	value := NewJsonValue(marshal)
	return value.Get(paths[1:]...)
}

func (m *mapValue) Slice() []Value {
	value := newErrorValue(errors.New("failed to as slice"))
	return value.Slice()
}

func (m *mapValue) String() string {
	ret, _ := json.Marshal(m.value)
	return string(ret)
}

func (m *mapValue) Err() error {
	return nil
}
