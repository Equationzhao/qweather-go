//go:build jsoniter

package json

import "github.com/json-iterator/go"

func Unmarshal(data []byte, v interface{}) error {
	return jsoniter.Unmarshal(data, v)
}

func Marshal(v interface{}) ([]byte, error) {
	return jsoniter.Marshal(v)
}
