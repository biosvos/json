package json

import (
	"encoding/json"
	"github.com/pkg/errors"
)

var _ Value = &mapValue{}

type mapValue struct {
	value map[string]any
}

func (m *mapValue) List(paths ...string) ([]Value, error) {
	return slice(m, paths...)
}

func newMapValue(value map[string]any) *mapValue {
	return &mapValue{
		value: value,
	}
}

func (m *mapValue) Get(paths ...string) (Value, error) {
	if len(paths) == 0 {
		return m, nil
	}
	v, ok := m.value[paths[0]]
	if !ok {
		return nil, errors.Errorf("not found path %v", paths[0])
	}
	marshal, _ := json.Marshal(v)
	value, _ := NewJsonValue(marshal)
	return value.Get(paths[1:]...)
}

func (m *mapValue) slice() ([]Value, error) {
	return nil, errors.New("failed to as slice")
}

func (m *mapValue) String() string {
	return string(m.Bytes())
}

func (m *mapValue) Bytes() []byte {
	ret, _ := json.Marshal(m.value)
	return ret
}
