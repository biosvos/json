package json

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strconv"
)

func Int(str string) (int64, error) {
	ret, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return ret, nil
}

func Bool(str string) (bool, error) {
	ret, err := strconv.ParseBool(str)
	if err != nil {
		return false, errors.WithStack(err)
	}
	return ret, nil
}

func Float(str string) (float64, error) {
	ret, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return 0, errors.WithStack(err)
	}
	return ret, nil
}

func Copy(dst any, src string) error {
	err := json.Unmarshal([]byte(src), dst)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
