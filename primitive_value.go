package json

import (
	"github.com/pkg/errors"
	"strconv"
	"strings"
)

var _ Value = &PrimitiveValue{}

type PrimitiveValue struct {
	value []byte
}

func (p *PrimitiveValue) Get(paths ...string) Value {
	if len(paths) == 0 {
		return p
	}
	return p.get(paths[0])
}

func NewPrimitiveValue(value []byte) *PrimitiveValue {
	return &PrimitiveValue{value: value}
}

func (p *PrimitiveValue) get(paths string) Value {
	return NewErrorValue(errors.Errorf("failed to get %v", paths))
}

func (p *PrimitiveValue) AsSlice() []Value {
	return nil
}

func (p *PrimitiveValue) String() (string, error) {
	return strings.Trim(string(p.value), "\""), nil
}

func (p *PrimitiveValue) Int() (int64, error) {
	str, _ := p.String()
	ret, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return ret, nil
}

func (p *PrimitiveValue) Bool() (bool, error) {
	str, _ := p.String()
	ret, err := strconv.ParseBool(str)
	if err != nil {
		return false, err
	}
	return ret, nil
}

func (p *PrimitiveValue) Bytes() []byte {
	return p.value
}
