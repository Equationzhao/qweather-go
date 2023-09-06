//go:build !jsoniter && !sonic

package json

import "encoding/json"

var Unmarshal = json.Unmarshal

var Marshal = json.Marshal
