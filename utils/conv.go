package utils

import "encoding/json"

func Conv(v interface{}) interface{} {
	switch v.(type) {
	case json.Number:
		return v.(json.Number).String()
	default:
		return v
	}
}

func Conva(v []interface{}) []interface{} {
	for i, attr := range v {
		v[i] = Conv(attr)
	}
	return v
}
