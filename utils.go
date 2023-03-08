package json

import "encoding/json"

func convert[T any](value any) (T, error) {
	var ret T
	marshal, err := json.Marshal(value)
	if err != nil {
		return ret, err
	}

	err = json.Unmarshal(marshal, &ret)
	if err != nil {
		return ret, err
	}
	return ret, nil
}
