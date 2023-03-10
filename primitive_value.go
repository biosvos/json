package json

import (
	"github.com/pkg/errors"
)

var _ Value = &primitiveValue{}

type primitiveValue struct {
	value []byte
}

func newPrimitiveValue(value []byte) *primitiveValue {
	return &primitiveValue{value: value}
}

func (p *primitiveValue) Get(paths ...string) Value {
	if len(paths) == 0 {
		return p
	}
	return newErrorValue(errors.Errorf("failed to get %v", paths[0]))
}

func (p *primitiveValue) Slice() []Value {
	value := newErrorValue(errors.New("failed to as slice"))
	return value.Slice()
}

func (p *primitiveValue) String() string {
	str := string(p.value)
	if len(str) < 3 {
		return str
	}
	if str[0] == '"' && str[len(str)-1] == '"' {
		return str[1 : len(str)-1]
	}
	return str
}

func (p *primitiveValue) Err() error {
	return nil
}
