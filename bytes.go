package json

import (
	"encoding/json"
	"github.com/pkg/errors"
)

func toMap(bytes []byte) (map[string]any, error) {
	m := map[string]any{}
	err := json.Unmarshal(bytes, &m)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return m, nil
}

func toSlice(bytes []byte) ([]any, error) {
	var ret []any
	err := json.Unmarshal(bytes, &ret)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return ret, nil
}
