package json

var _ Value = &ErrorValue{}

type ErrorValue struct {
	value error
}

func NewErrorValue(value error) *ErrorValue {
	return &ErrorValue{value: value}
}

func (e *ErrorValue) Get(_ ...string) Value {
	return e
}

func (e *ErrorValue) AsSlice() []Value {
	return nil
}

func (e *ErrorValue) String() string {
	return ""
}

func (e *ErrorValue) Int() (int64, error) {
	return 0, e.value
}

func (e *ErrorValue) Bool() (bool, error) {
	return false, e.value
}

func (e *ErrorValue) Bytes() []byte {
	return nil
}
