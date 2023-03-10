package json

import "strconv"

func Int(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}

func Bool(str string) (bool, error) {
	return strconv.ParseBool(str)
}

func Float(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}
