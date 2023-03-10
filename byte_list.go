package json

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type byteList []byte

func (b byteList) toMap() (map[string]any, error) {
	m := map[string]any{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return m, nil
}

func (b byteList) toSlice() ([]any, error) {
	var ret []any
	err := json.Unmarshal(b, &ret)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ret, nil
}
