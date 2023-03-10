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

func (p *primitiveValue) Get(paths ...string) (Value, error) {
	if len(paths) == 0 {
		return p, nil
	}
	return nil, errors.Errorf("failed to get %v", paths[0])
}

func (p *primitiveValue) AsSlice() ([]Value, error) {
	return nil, errors.New("failed to as slice")
}

func (p *primitiveValue) String() string {
	str := string(p.Bytes())
	if len(str) < 3 {
		return str
	}
	if str[0] == '"' && str[len(str)-1] == '"' {
		return str[1 : len(str)-1]
	}
	return str
}

func (p *primitiveValue) Bytes() []byte {
	return p.value
}
