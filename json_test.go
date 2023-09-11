package qweather

import (
	"bytes"
	stdjson "encoding/json"
	"strings"
	"testing"

	"github.com/Equationzhao/qweather-go/internal/json"
)

func _myMarshal(v any) ([]byte, error) {
	bb := bytes.Buffer{}
	m, err := stdjson.Marshal(v)
	if err != nil {
		return nil, err
	}
	bb.WriteString("my Marshal")
	bb.Write(m)
	return bb.Bytes(), nil
}

func _myUnmarshal(b []byte, v any) error {
	b, _ = bytes.CutPrefix(b, []byte("my Marshal"))
	return stdjson.Unmarshal(b, v)
}

func TestSetJsonMarshal(t *testing.T) {
	SetJsonMarshal(_myMarshal)
	v := struct {
		A string `json:"a"`
	}{
		A: "test",
	}
	res, err := json.Marshal(v)
	if err != nil {
		t.Fatal(err)
	}
	if strings.HasPrefix(string(res), "my Marshal") {
		t.Log("success")
	} else {
		t.Fatal("fatal")
	}
}

func TestSetJsonUnmarshal(t *testing.T) {
	SetJsonUnmarshal(_myUnmarshal)
	jsonString := "my Marshal{\"a\":\"test\"}"
	v := struct {
		A string `json:"a"`
	}{}
	err := json.Unmarshal([]byte(jsonString), &v)
	if err != nil {
		t.Fatal(err)
	}
	if v.A == "test" {
		t.Log("success")
	} else {
		t.Fatal(v)
	}
}
