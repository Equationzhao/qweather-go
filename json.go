//go:build !sonic && !jsoniter

package qweather

import "github.com/Equationzhao/qweather-go/internal/json"

func SetJsonMarshal(j func(v any) ([]byte, error)) {
	json.Marshal = j
}

func SetJsonUnmarshal(j func(data []byte, v any) error) {
	json.Unmarshal = j
}
