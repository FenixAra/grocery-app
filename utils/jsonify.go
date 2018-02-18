package utils

import "encoding/json"

func Jsonify(v interface{}) string {
	data, _ := json.Marshal(v)
	return string(data)
}
