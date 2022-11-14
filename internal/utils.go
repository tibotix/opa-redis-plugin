package internal

import "encoding/json"

func conv(v interface{}) interface{} {
	switch v.(type) {
	case json.Number:
		return v.(json.Number).String()
	default:
		return v
	}
}

func conva(v []interface{}) []interface{} {
	for i, attr := range v {
		v[i] = conv(attr)
	}
	return v
}
