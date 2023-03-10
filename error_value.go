package json

var _ Value = &errorValue{}

type errorValue struct {
	error error
}

func newErrorValue(error error) *errorValue {
	return &errorValue{error: error}
}

func (e *errorValue) Get(_ ...string) Value {
	return e
}

func (e *errorValue) Slice() []Value {
	return []Value{e}
}

func (e *errorValue) String() string {
	return ""
}

func (e *errorValue) Err() error {
	return e.error
}
