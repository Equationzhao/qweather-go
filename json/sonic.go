//go:build sonic

package json

import "github.com/bytedance/sonic"

func Unmarshal(data []byte, v interface{}) error {
	return sonic.Unmarshal(data, v)
}

func Marshal(v interface{}) ([]byte, error) {
	return sonic.Marshal(v)
}
